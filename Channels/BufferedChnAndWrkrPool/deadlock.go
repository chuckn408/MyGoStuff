package main

import (
	"fmt"
)

func main() {
	ch := make(chan string, 2)
	ch <- "naveen"
	ch <- "paul"
	ch <- "steve"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

/*
 go run .\deadlock.go
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()
*/