package main

import (
	"fmt"
)

func main() {
	str := "Gopher"
	fmt.Printf("Original String: %s\n", string(str))
	fmt.Printf("Reversed String: ")
	for _, v := range str {
		defer fmt.Printf("%c", v)
	}
}