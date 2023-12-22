// server1 是一个迷你回声服务器
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", echo)
	http.ListenAndServe(":2333", nil)
}

func echo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
