package main

import (
	"bufio"
	"os"
)

func balancedStringSplit(s string) int {
	var cnt, n int
	for i := 0; i < len(s); i++ {
		if s[i] == 'R' {
			n++
		} else if s[i] == 'L' {
			n--
		}
		if n == 0 {
			cnt++
		}
	}
	return cnt
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		println(balancedStringSplit(input.Text()))
	}
}
