package main

import "fmt"

func main() {
	n := make(chan int)
	s := make(chan int)

	go in(n)
	go out(s, n)

	printer(s)
}

func printer(s chan int) {
	for s := range s {
		fmt.Println(s)
	}
}

func out(s chan int, n chan int) {
	for x := range n {
		s <- x * x
	}
	close(s)
}

func in(n chan int) {
	for i := 0; i <= 100; i++ {
		n <- i
	}
	close(n)
}
