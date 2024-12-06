package main

import (
	"fmt"
)

func main() {
	a := func() {
		fmt.Println("hello world first class func")
	}
	
	a()
	fmt.Printf("%T", a) // this prints what a's "T"ype is
}