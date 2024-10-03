package main

import (
	"SomeProject/greetings"
	"fmt"
	"log"
)

func main() {
	// Set prop of predef log, incl log entry prefix and flg to dis prnt time, src and LN
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names.
	names := []string{"Gladys", "Samantha", "Darrin"}

	// Request greeting msgs for names
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	// If no err rtrnd, prnt ret map o msg to cnsl
	fmt.Println(messages)
}
