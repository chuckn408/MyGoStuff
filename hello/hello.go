package main

import (
	"GreetAgain/greetings"
	"fmt"
)

func main() {
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
