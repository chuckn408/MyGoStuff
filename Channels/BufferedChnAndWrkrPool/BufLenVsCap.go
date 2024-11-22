package main

import (
	"fmt"
)

func main() {
	ch := make(chan string, 3)
	ch <- "naveen"
	ch <- "pual"
	fmt.Println("capacity is", cap(ch))
	fmt.Println("length is", len(ch))
	fmt.Println("read val is", <-ch)
	fmt.Println("new len is ", len(ch))
}