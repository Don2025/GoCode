package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/jian-sheng-zi-ii-lcof/
// ------------------------剑指 Offer I Problem 14-II------------------------
func cuttingRope(n int) int {
	if n < 4 {
		return n - 1
	}
	const mod = 1000000007
	ans := 1
	for n > 4 {
		ans = ans * 3 % mod
		n -= 3
	}
	return ans * n % mod
}

// ------------------------剑指 Offer I Problem 14-II------------------------
/*
 * https://leetcode.cn/problems/jian-sheng-zi-ii-lcof/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.8MB 在所有Go提交中击败了72.85%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", cuttingRope(n))
	}
}
