package main

import "fmt"

func main() {

	ch := make(chan string, 3)
	ch <- "1"
	ch <- "2"
	fmt.Println(len(ch), cap(ch))

	ch <- "3"
	//ch <- "4" //fatal error: all goroutines are asleep - deadlock!

	fmt.Println(len(ch), cap(ch))
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	//fmt.Println(<-ch) //fatal error: all goroutines are asleep - deadlock!
}
