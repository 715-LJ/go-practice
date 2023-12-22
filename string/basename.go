package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(basename("a"))
	fmt.Println(basename("b.go"))
	fmt.Println(basename("a/b/c.go"))
	fmt.Println(basename("a/b.c.go"))
}

func basename(s string) string {
	slash := strings.LastIndex(s, "/") // 获取/下标，不存在返回-1
	s = s[slash+1:]
	fmt.Printf("文件名称 -- %s\n", s)
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}
