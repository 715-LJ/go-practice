package main

import (
	"fmt"
	"sync"
)

//sync.WaitGroup 提供了三个方法：
//Add(delta int)：用于向 WaitGroup 添加计数器，delta 为正表示增加计数器的值，delta 为负表示减少计数器的值。每次调用 Add 方法都会增加或减少计数器的值。
//Done()：表示一个 goroutine 已经完成了任务，会减少 WaitGroup 的计数器值，相当于调用了 Add(-1)。
//Wait()：阻塞当前 goroutine，直到 WaitGroup 的计数器值变为 0。一旦计数器值为 0，Wait 方法就会立即返回，程序继续执行后续的操作。

func main() {

	fmt.Println("start...")
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			fmt.Println(i)
		}()
	}
	wg.Wait()

	fmt.Println("end...")
}
