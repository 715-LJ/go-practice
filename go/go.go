package main

import "fmt"

var balance int

func Save(num int) {
	balance += num
}

func main() {

	for i := 1; i <= 10; i++ {
		go func() {
			Save(10)
		}()
		fmt.Println(i, balance)
	}
}
