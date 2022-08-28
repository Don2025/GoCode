package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/maximum-product-of-two-elements-in-an-array/
//------------------------Leetcode Problem 1464------------------------
func maxProduct(nums []int) int {
	a, b := nums[0], nums[1]
	if a < b {
		a, b = b, a
	}
	for i := 2; i < len(nums); i++ {
		if nums[i] > a {
			a, b = nums[i], a
		} else if nums[i] > b {
			b = nums[i]
		}
	}
	return (a - 1) * (b - 1)
}

//------------------------Leetcode Problem 1464------------------------
/*
 * https://leetcode.cn/problems/maximum-product-of-two-elements-in-an-array/
 * 执行用时：0ms 在所有Go提交中击败了26.79%的用户
 * 占用内存：2.8MB 在所有Go提交中击败了80.36%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		Printf("Output: %d\n", maxProduct(nums))
	}
}
