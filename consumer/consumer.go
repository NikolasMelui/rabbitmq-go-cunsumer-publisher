package main

import (
	"os"
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

}
