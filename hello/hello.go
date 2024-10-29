package main

import (
	"fmt"
	"greetings/greetings"
)

func main() {
	message := greetings.Hello("Glayds")
	fmt.Println(message)
}
