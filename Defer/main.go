package main

import (
	"fmt"
	"time"
)

func totalTime(start time.Time) {
	fmt.Printf("Total time taken %f seconds", time.Since(start).Seconds())
}

func test() {
	start := time.Now()
	defer totalTime(start)
	time.Sleep(2 * time.Second)
	fmt.Println("Sleep complete")
}

func main() {
	test()
}

/*
The above is a simple program which illustrates the use of defer. In the 
above program, defer is used to find out the total time taken for the 
execution of the test() function. The start time of the test() function
 execution is passed as argument to defer totalTime(start) in line no. 14.
 This defer call is executed just before test() returns. totalTime prints 
 the difference between start and the current time using time.Since in line 
 no. 9. To simulate some computation happening in test(), a 2 second sleep
 is added in line no. 15.
 */