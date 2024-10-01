package main

import (
	"SomeProject/greetings"
	"fmt"
	"log"
)

func main() {
	// Set properties of the predefined Logger, incld entry pref and flag to disbl prnt the time, src and ln
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting message.
	message, err := greetings.Hello("Gladys")
	// If an err was retn, prnt to cnsl and exit
	if err != nil {
		log.Fatal(err)
	}

	// if no err, prnt return msg to cnsl
	fmt.Println(message)
}
