package main

import (
	"bufio"
	"os"
	"strconv"
)

func isPowerOfThree(n int) bool {
	if n < 3 {
		return n == 1
	}
	if n%3 == 0 {
		return isPowerOfThree(n / 3)
	}
	return false
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		println(isPowerOfThree(n))
	}
}

/*
 * 执行用时：20ms 在所有Go提交中击败了91.48%的用户
 * 占用内存：6.4MB 在所有Go提交中击败了70.33%的用户
**/
