package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greet for named person
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return name, errors.New("empty name")
	}
	// Create a message using a random format.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Hellos return a map that associates each named w/ a msg
func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	// Loop through the received slice of names, calliing the Hel func to get msg for ea name
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// In map, assoc the retrieve msg w/ name.
		messages[name] = message
	}
	return messages, nil
}

// randomFormat returns one of a set of greet msgs. Retrn msg is rand
func randomFormat() string {
	// A slice of msg frmts
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// Rtrn one of the msg frmt sel at rndm.
	return formats[rand.Intn(len(formats))]
}
