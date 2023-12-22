package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		if strings.TrimSpace(input.Text()) == "#" {
			break
		}
		counts[input.Text()]++
	}
}
