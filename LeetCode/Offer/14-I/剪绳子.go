package main

import (
	"bufio"
	. "fmt"
	"math"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/jian-sheng-zi-lcof/
// ------------------------剑指 Offer I Problem 14-I------------------------
func cuttingRope(n int) int {
	if n <= 3 {
		return n - 1
	}
	return int(math.Pow(3, float64(n/3)) * 4.0 / float64(4-n%3))
}

// ------------------------剑指 Offer I Problem 14-I------------------------
/*
 * https://leetcode.cn/problems/jian-sheng-zi-lcof/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了41.73%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", cuttingRope(n))
	}
}
