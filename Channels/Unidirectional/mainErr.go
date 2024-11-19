package main

import "fmt"

func sendData(sendch chan<- int) {
	sendch <- 10
}

func main() {
	sendch := make(chan<- int)
	go sendData(sendch)
	fmt.Println(<-sendch)
}

/*
./prog.go:12:14: invalid operation: <-sendch (receive from send-only type chan<- int)
*/