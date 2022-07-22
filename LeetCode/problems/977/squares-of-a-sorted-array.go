package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
)

// https://leetcode.cn/problems/squares-of-a-sorted-array/
//------------------------Leetcode Problem 977------------------------
func sortedSquares(nums []int) []int {
	ans := make([]int, len(nums))
	l, r := 0, len(nums)-1
	for i := len(nums) - 1; l <= r; i-- {
		if abs(nums[l]) > abs(nums[r]) {
			ans[i] = pow2(nums[l])
			l++
		} else {
			ans[i] = pow2(nums[r])
			r--
		}
	}
	return ans
}

func abs(n int) int {
	if n < 0 {
		n = -n
	}
	return n
}

func pow2(n int) int {
	return n * n
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := stringArrayToIntArray(strings.Fields(scanner.Text()))
		Printf("Output: %v\n", sortedSquares(nums))
	}
}

//------------------------Leetcode Problem 977------------------------
/*
 * https://leetcode.cn/problems/squares-of-a-sorted-array/
 * 执行用时：28ms 在所有Go提交中击败了80.78%的用户
 * 占用内存：6.8MB 在所有Go提交中击败了64.13%的用户
**/
