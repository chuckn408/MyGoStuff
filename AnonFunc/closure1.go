package main

import (
	"fmt"
)

func main() {
	a := 5
	func() {
		fmt.Println("a =", a)
	}() //again, cant move parenthesis or itll break
}