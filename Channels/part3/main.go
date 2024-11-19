package main

import (
	"fmt"
	"time"
)

func hello(done chan bool) {
	fmt.Println("hello go rout is slep")
	time.Sleep(4 * time.Second)
	fmt.Println("hello go rou awake and write done")
	done <- true
}

func main() {
	done := make(chan bool)
	fmt.Println("Main goin to call hello gorout")
	go hello(done)
	<-done
	fmt.Println("Main received data")
}