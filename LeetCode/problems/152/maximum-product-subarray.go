package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/maximum-product-subarray/
//------------------------Leetcode Problem 152------------------------
func maxProduct(nums []int) int {
	maxF, minF, ans := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		mx, mn := maxF, minF
		maxF = max(mx*nums[i], max(nums[i], mn*nums[i]))
		minF = min(mx*nums[i], min(nums[i], mn*nums[i]))
		ans = max(ans, maxF)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//------------------------Leetcode Problem 152------------------------
/*
 * https://leetcode.cn/problems/maximum-product-subarray/
 * 执行用时：4ms 在所有Go提交中击败了76.43%的用户
 * 占用内存：3.3MB 在所有Go提交中击败了33.64%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", maxProduct(nums))
	}
}
