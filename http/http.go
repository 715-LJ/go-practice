package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	url := "https://10.41.56.241:9443"

	if !strings.HasPrefix(url, "https://") {
		url = "https://" + url
	}

	fmt.Println(url)
	resp, err := http.Get(url)

	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %v\n", err)
		os.Exit(1)
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stdout, "err: %v\n", err)
	}

	fmt.Printf("%s\n", b) // []byte 可使用 %s 直接输出
}
