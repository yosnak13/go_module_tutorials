package main

import (
	"fmt"
	"log"

	"exmaple.com/greetings"
)

func main() {

	// Set Properties of the predefined Logger, including the log entry prefix and flag to disable printing the time,
	// source file, and line number/
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Gladys")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
