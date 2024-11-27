package main

func main() {
	ch := make(chan string)
	select {
		case <-ch:
	}
}

// result:: fatal error: all goroutines are asleep - deadlock!