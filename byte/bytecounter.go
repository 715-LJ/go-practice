package main

import "fmt"

type ByteCounter int

func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c)

	c = 0
	fmt.Fprintf(&c, "I am lijia")
	fmt.Println(c)
}

func (c *ByteCounter) Write(p []byte) (n int, err error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}
