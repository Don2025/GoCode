package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

func isIdealPermutation(nums []int) bool {
	n := len(nums)
	minSuf := nums[n-1]
	for i := n - 2; i > 0; i-- {
		if nums[i-1] > minSuf {
			return false
		}
		minSuf = min(minSuf, nums[i])
	}
	return true
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//------------------------Leetcode Problem 775------------------------
/*
 * https://leetcode.cn/problems/global-and-local-inversions/
 * 执行用时：100ms 在所有Go提交中击败了55.56%的用户
 * 占用内存：8.9MB 在所有Go提交中击败了94.44%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", isIdealPermutation(nums))
	}
}
