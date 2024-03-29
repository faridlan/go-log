package test

import (
	"log"
	"os"
	"testing"
)

func TestHelloWorld(t *testing.T) {

	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)
	log.Println("Hello World")
}
