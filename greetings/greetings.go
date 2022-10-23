package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	// If no name was given, return an error with message.
	if name == "" {
		return "", errors.New("empty name")
	}

	// If a name was received, return an value that embeds tha name in a greeting message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
