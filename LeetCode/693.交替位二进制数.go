package main

import (
	"bufio"
	"os"
	"strconv"
)

func hasAlternatingBits(n int) bool {
	n = n ^ (n >> 1)
	return (n & (n + 1)) == 0
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		println(hasAlternatingBits(n))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.8MB 在所有Go提交中击败了85.33%的用户
**/
