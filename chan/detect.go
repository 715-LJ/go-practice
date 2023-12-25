// TCP 服务器
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("listen ...")

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("okkkk")
		handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	input := bufio.NewScanner(conn)
	ch := make(chan bool)
	go detect(conn, ch)
	for input.Scan() {
		ch <- true
		go echo(conn, input.Text(), 1*time.Second)
	}
}

func echo(conn net.Conn, text string, duration time.Duration) {
	fmt.Fprintln(conn, "\t", strings.ToUpper(text))
	time.Sleep(duration)

	fmt.Fprintln(conn, "\t", text)
	time.Sleep(duration)
	fmt.Fprintln(conn, "\t", strings.ToLower(text))
	time.Sleep(duration)
}

func detect(conn net.Conn, ch chan bool) {
	maxIdle := 10
	idleSeconds := 0
	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ticker.C:
			idleSeconds++
			if idleSeconds > maxIdle { //释放链接
				msg := conn.RemoteAddr().String() + " 10s has no request, kicked out"
				fmt.Println(msg)
				fmt.Fprintf(conn, msg)
				conn.Close()
				close(ch)
				ticker.Stop()
				return
			}
		case <-ch:
			idleSeconds = 0 //重制计数器
		}
	}
}
