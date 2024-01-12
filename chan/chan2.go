package main

import "fmt"

func main() {
	counter := make(chan int)
	square := make(chan int)

	go func() {
		for x := 0; x <= 5; x++ {
			counter <- x
		}
		close(counter)
	}()

	go func() {
		for x := range counter {
			square <- x * x
		}
		close(square)
	}()

	for res := range square {
		fmt.Println(res)
	}
}
