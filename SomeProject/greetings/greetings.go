package greetings

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) string { //functions starting with 'Captials' are 'exported names'.
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("hi, %v. Welcome!", name) // := is a init/declr shortcut
	return message
}
