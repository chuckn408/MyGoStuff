package main

import "fmt"

func main() {
	var a chan int
	if a == nil {
		fmt.Println("channel a is nil, goin got define it")
		a = make(chan int)  //chanel 'a'
		fmt.Printf("Type of a is %T", a)
		// shorthand declaration --> a := make(chan int)
	}
}