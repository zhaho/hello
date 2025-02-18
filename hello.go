package main

import (
	"fmt"
	"log"

	"github.com/zhaho/greetings"
)

func main() {
	log.SetPrefix("greetings ")
	log.SetFlags(0)

	// Get a greeting message and print it.
	message, err := greetings.Hello("Kenneth")

	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println(message)
}