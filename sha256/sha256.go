package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {

	sli1 := []byte("lijia")
	sha1 := sha256.Sum256(sli1)
	fmt.Println(sha1)
}
