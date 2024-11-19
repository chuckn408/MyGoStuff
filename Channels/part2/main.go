package main

import (
	"fmt"
)

func hello(done chan bool) {
	fmt.Println("hello world goroutine")
	done <- true
}

func main() {
	done := make(chan bool)
	go hello(done)
	<-done  //until some Gorou wrts 2 done ch,cntrl waits
	fmt.Println("main function")
}