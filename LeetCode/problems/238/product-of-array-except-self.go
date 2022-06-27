package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/product-of-array-except-self/
//------------------------Leetcode Problem 238------------------------
func productExceptSelf(nums []int) []int {
	ans := make([]int, len(nums))
	for i := range ans {
		ans[i] = 1
	}
	left, right := 1, 1
	for i := 0; i < len(nums); i++ {
		ans[i] *= left
		left *= nums[i]
		ans[len(nums)-i-1] *= right
		right *= nums[len(nums)-i-1]
	}
	return ans
}

//------------------------Leetcode Problem 238------------------------
/*
 * https://leetcode.cn/problems/product-of-array-except-self/
 * 执行用时：32ms 在所有Go提交中击败了81.68%的用户
 * 占用内存：7.6MB 在所有Go提交中击败了52.90%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", productExceptSelf(nums))
	}
}
