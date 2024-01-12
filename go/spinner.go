package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond)
}

func spinner(duration time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(duration)
		}
	}
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
