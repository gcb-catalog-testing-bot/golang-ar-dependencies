package main

import (
	"fmt"
	"log"

	greetings "gcb-catalog-dev.git.corp.google.com/gcb-catalog-testing-bot/golang-ar-private-module"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Catalog")

	if message == "" {
		log.Fatal("error: Message returned empty")
	}
	fmt.Println(message)
}
