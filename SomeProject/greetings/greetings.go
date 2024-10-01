package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name was givern, return an error with a message.
	if name == "" {
		return name, errors.New("empty name")
	}
	// Create a message using a random format.
	message := fmt.Sprint(randomFormat(), name)
	return message, nil
}

// reandomFormat returns one of the set of greeting messages. The returned msg is at random
func randomFormat() string {
	// A slice of msg formats
	formats := []string{
		"hi, %v. Welcom!",
		"great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// Return a rand sel msg fmt by spec a rnd ind for the slc of fmts.
	return formats[rand.Intn(len(formats))]
}
