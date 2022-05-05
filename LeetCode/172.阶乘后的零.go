package main

import (
	"bufio"
	"os"
	"strconv"
)

func trailingZeroes(n int) int {
	var ans int
	for n > 0 {
		n /= 5
		ans += n
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		println(trailingZeroes(n))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.8MB 在所有Go提交中击败了100.00%的用户
**/
