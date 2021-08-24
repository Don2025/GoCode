package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func generateTheString(n int) string {
	if n&1 == 0 {
		return strings.Repeat("s", n-1) + "b"
	} else {
		return strings.Repeat("s", n)
	}
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		println(generateTheString(n))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.1MB 在所有Go提交中击败了81.25%的用户
**/
