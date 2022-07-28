package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/w3tCBm/
// ------------------------剑指 Offer II Problem 3------------------------
func countBits(n int) []int {
	ans := make([]int, n+1)
	for i := 0; i <= n; i++ {
		ans[i] = ans[i>>1] + (i & 1)
	}
	return ans
}

// ------------------------剑指 Offer II Problem 3------------------------
/*
 * https://leetcode.cn/problems/w3tCBm/
 * 执行用时：4ms 在所有Go提交中击败了83.47%的用户
 * 占用内存：4.4MB 在所有Go提交中击败了100.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		Printf("%v\n", countBits(n))
	}
}
