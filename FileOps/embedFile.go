package main

import (
	_ "embed"
	"fmt"
)

//go:embed test.txt
// remember, the above is not a normal comment; it goes to the compiler

var contents []byte

func main() {
	fmt.Println("contents of file: ", string(contents))
}