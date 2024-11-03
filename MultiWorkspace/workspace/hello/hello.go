package main

import (
	"fmt"
	"golang.org/x/example/hello/reverse"
)

func main() {
	//	fmt.Println(reverse.String("Hello"), reverse.Int(24601))
	//  this doesnt seem to work for the Int type
	fmt.Println(reverse.String("Hello"), reverse.String("24601"))
}
