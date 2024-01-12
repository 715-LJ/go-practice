package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"/Users/lijia/Desktop/LiJia/code/go-practice"}
	}

	fmt.Println(roots)

	fileSizes := make(chan int64)
	go func() {
		for _, dir := range roots {
			walkDir(dir, fileSizes)
		}
		close(fileSizes)
	}()
}

func walkDir(dir string, sizes chan int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := path.Join(dir, entry.Name())
			walkDir(subdir, sizes)
		} else {
			sizes <- entry.Size()
		}
	}
}

func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return entries
}
