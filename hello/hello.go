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

	// A slice of names
	names := []string{"Gladys", "Samantha", "Darrin"}

	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}
