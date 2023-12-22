package main

import "fmt"

func main() {

	sli1 := []int{1, 2, 3, 4, 5}
	reverse(sli1)

	fmt.Println(sli1)
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}