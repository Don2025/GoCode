package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/rotate-function/
//------------------------Leetcode Problem 396------------------------
func maxRotateFunction(nums []int) int {
	var f, sum int
	for i := range nums {
		sum += nums[i]
		f += i * nums[i]
	}
	ans, n := f, len(nums)
	for i := 1; i < n; i++ {
		f += sum - n*nums[n-i]
		ans = max(ans, f)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//------------------------Leetcode Problem 396------------------------
/*
 * https://leetcode.cn/problems/rotate-function/
 * 执行用时：120ms 在所有Go提交中击败了91.43%的用户
 * 占用内存：7.9MB 在所有Go提交中击败了65.71%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", maxRotateFunction(nums))
	}
}
