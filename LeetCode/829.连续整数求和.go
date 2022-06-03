package main

import (
	"bufio"
	"os"
	"strconv"
)

func consecutiveNumbersSum(n int) int {
	var ans int
	for i := 1; n > 0; i++ {
		if n%i == 0 {
			ans++
		}
		n -= i
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		println(consecutiveNumbersSum(n))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.8MB 在所有Go提交中击败了78.26%的用户
**/
