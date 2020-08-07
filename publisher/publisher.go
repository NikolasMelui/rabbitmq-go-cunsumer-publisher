package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/streadway/amqp"
)

var rabbitHost = os.Getenv("RABBIT_HOST")
var rabbitPort = os.Getenv("RABBIT_PORT")
var rabbitUser = os.Getenv("RABBIT_USERNAME")
var rabbitPassword = os.Getenv("RABBIT_PASSWORD")

func main() {
	router := httprouter.New()
	router.POST("/publish/:message", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		message := p.ByName("message")
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
				Body:        []byte(message),
			},
		)
		if err != nil {
			log.Fatalf("%s: %s", "Failed to publish a message", err)
		}

	})
	fmt.Println("Running...")
	log.Fatal(http.ListenAndServe(":8081", router))
}
