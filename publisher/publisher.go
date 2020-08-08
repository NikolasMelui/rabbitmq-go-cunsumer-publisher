package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

var rabbitHost = os.Getenv("RABBIT_HOST")
var rabbitPort = os.Getenv("RABBIT_PORT")
var rabbitUser = os.Getenv("RABBIT_USERNAME")
var rabbitPassword = os.Getenv("RABBIT_PASSWORD")

// Data ...
type Data struct {
	Lang string `json:"lang"`
	Code string `json:"code"`
}

func main() {
	router := httprouter.New()
	router.POST("/publish", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatalf("%s: %s", "Failed to read the body of the request", err)
		}

		var data Data
		json.Unmarshal(body, &data)

		message := fmt.Sprint(data)
		fmt.Println("Received message:" + message)

		amqpDialAddr := fmt.Sprintf("amqp://%s:%s@%s:%s/", rabbitUser, rabbitPassword, rabbitHost, rabbitPort)

		conn, err := amqp.Dial(amqpDialAddr)
		if err != nil {
			log.Fatalf("%s: %s", "Failed to connect to RabbitMQ", err)
		}
		defer conn.Close()

		ch, err := conn.Channel()
		if err != nil {
			log.Fatalf("%s: %s", "Failed to open a channel", err)
		}
		defer ch.Close()

		q, err := ch.QueueDeclare(
			"publisher",
			false,
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			log.Fatalf("%s: %s", "Failed to declare a queue", err)
		}

		err = ch.Publish(
			"",
			q.Name,
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        body,
			},
		)
		if err != nil {
			log.Fatalf("%s: %s", "Failed to publish a message", err)
		}

	})
	fmt.Println("Running...")
	log.Fatal(http.ListenAndServe(":8081", router))
}
