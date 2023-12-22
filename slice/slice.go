package main

import "fmt"

func main() {
	sli1 := []byte("a")
	fmt.Println(sli1)

	sli2 := []string{"lijia01", "lijia02", "lijia03"}
	fmt.Println(sli2)
	fmt.Println(sli2[:len(sli2)-1])
}
