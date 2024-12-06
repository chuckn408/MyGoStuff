package main

import (
	"fmt"
)

func main() {
	func() {
		fmt.Println("hello from no var named func")
	}() // cant move parenthesis or itll break
}