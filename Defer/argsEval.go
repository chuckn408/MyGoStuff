package main

import (
	"fmt"
)

func displayValue(a int) {
	fmt.Println("value of a in deferred fucn", a)
}

func main() {
	a := 5
	defer displayValue(a)
	a = 10
	fmt.Println("value of a before deferred function all", a)
}

/*
In the program above a initially has a value of 5 in line no. 11.
 When the defer statement is executed in line no. 12, the value 
 of a is 5 and hence this will be the argument to the display
 Value function which is deferred. We change the value of a to 10 
 in line no. 13. The next line prints the value of a. This 
 program outputs,
*/