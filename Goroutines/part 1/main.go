package main

import (
"fmt"
)

func hello() {
	fmt.Println("hello world goroutine")
}

func main() {
	go hello()
	fmt.Println("main function")
}