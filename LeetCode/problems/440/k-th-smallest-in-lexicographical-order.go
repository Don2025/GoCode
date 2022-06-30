package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/k-th-smallest-in-lexicographical-order/
//------------------------Leetcode Problem 440------------------------
func findKthNumber(n int, k int) int {
	ans := 1
	var getSteps func(int, int) int
	getSteps = func(cur int, n int) int {
		first, last := cur, cur
		var steps int
		for first <= n {
			steps += min(last, n) - first + 1
			first *= 10
			last = last*10 + 9
		}
		return steps
	}
	k--
	for k > 0 {
		steps := getSteps(ans, n)
		if steps <= k {
			k -= steps
			ans++
		} else {
			ans *= 10
			k--
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//------------------------Leetcode Problem 440------------------------
/*
 * https://leetcode.cn/problems/k-th-smallest-in-lexicographical-order/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.8MB 在所有Go提交中击败了65.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		k, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", findKthNumber(n, k))
	}
}
