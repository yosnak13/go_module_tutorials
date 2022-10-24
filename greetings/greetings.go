package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {
	// If no name was given, return an error with message.
	if name == "" {
		return "", errors.New("empty name")
	}

	// Create a message using a random format.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
	}
	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}
