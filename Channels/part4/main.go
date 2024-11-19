package main

import (
	"fmt"
)

func calcSquares(number int, squareop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	squareop <- sum
}

func calcCubes(number int, cubeop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	cubeop <- sum
}

func main() {
	number := 589
	sqrch := make(chan int)
	cubech := make(chan int)
	go calcSquares(number, sqrch)
	go calcCubes(number, cubech)
	squares, cubes := <-sqrch, <-cubech
	fmt.Println("Final output", squares + cubes)
}

/*
The calcSquares function in line no. 7 calculates the sum of the squares
 of the individual digits of the number and sends it to the squareop 
 channel. Similarly the calcCubes function in line no. 17 calculates 
 the sum of cubes of the individual digits of the number and sends it 
 to the cubeop channel.

These two functions are run as separate Goroutines in line no. 31 and
 32 and each is passed a channel to write to as the parameter. The 
 main Goroutine waits for data from both these channels in line no.
 33. Once the data is received from both the channels, they are stored 
 in squares and cubes variables and the final output is computed and 
 printed. This program will print
 */