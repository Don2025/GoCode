package main

import (
	"bufio"
	"os"
	"strconv"
)

func binaryGap(n int) int {
	var ans int
	for i, j := 0, -1; n > 0; i++ {
		if n&1 == 1 {
			if j != -1 {
				ans = max(ans, i-j)
			}
			j = i
		}
		n >>= 1
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		println(binaryGap(n))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.8MB 在所有Go提交中击败了73.81%的用户
**/
