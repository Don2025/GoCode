package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/tvdfij/
// ------------------------剑指 Offer II Problem 12------------------------
func pivotIndex(nums []int) int {
	var sum, cur int
	for _, x := range nums {
		sum += x
	}
	for i := range nums {
		sum -= nums[i]
		if cur == sum {
			return i
		}
		cur += nums[i]
	}
	return -1
}

// ------------------------剑指 Offer II Problem 12------------------------
/*
 * https://leetcode.cn/problems/tvdfij/
 * 执行用时：24ms 在所有Go提交中击败了11.55%的用户
 * 占用内存：6MB 在所有Go提交中击败了72.52%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", pivotIndex(nums))
	}
}
