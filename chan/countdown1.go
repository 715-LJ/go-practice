package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("倒计时start...")
	tick := time.Tick(1 * time.Second)
	for {
		fmt.Println(time.Now())
		<-tick
	}
}
