package main

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	ch <- "from server 1"
}

func server2(ch chan string) {
	ch <- "form server 2"
}

func main() {
	output1 := make(chan string)
	output2 := make(chan string)
	go server1(output1)
	go server2(output2)
	time.Sleep(1 * time.Second)
	select {
		case s1 := <-output1:
			fmt.Println(s1)
		case s2 := <-output2:
			fmt.Println(s2)
	}
}