package main

import "fmt"

func main() {
	src := []int{1, 2, 3, 4, 5}
	dst := make([]int, 3)

	copy(dst, src)

	fmt.Printf("%v \n", dst)
}
