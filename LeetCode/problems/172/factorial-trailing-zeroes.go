package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/factorial-trailing-zeroes/
//------------------------Leetcode Problem 172------------------------
func trailingZeroes(n int) int {
	var ans int
	for n > 0 {
		n /= 5
		ans += n
	}
	return ans
}

//------------------------Leetcode Problem 172------------------------
/*
 * https://leetcode.cn/problems/factorial-trailing-zeroes/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.8MB 在所有Go提交中击败了100.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", trailingZeroes(n))
	}
}
