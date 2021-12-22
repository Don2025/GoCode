package main

import (
	"bufio"
	"os"
	"strings"
)

func repeatedStringMatch(a string, b string) int {
	times := len(b)/len(a) + 1
	s := a
	for i := 1; i <= times+1; i++ {
		if strings.Contains(s, b) {
			return i
		}
		s += a
	}
	return -1
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		arr := strings.Fields(input.Text())
		a, b := arr[0], arr[1]
		println(repeatedStringMatch(a, b))
	}
}

/*
 * 执行用时：16ms 在所有Go提交中击败了43.18%的用户
 * 占用内存：8.5MB 在所有Go提交中击败了34.09%的用户
**/
