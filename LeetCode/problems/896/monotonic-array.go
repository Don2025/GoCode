package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/monotonic-array/
//------------------------Leetcode Problem 896------------------------
func isMonotonic(nums []int) bool {
	n := len(nums)
	if n == 1 {
		return true
	}
	var a, b int
	for i := 0; i < n-1; i++ {
		if nums[i] < nums[i+1] {
			a = 1
		} else if nums[i] > nums[i+1] {
			b = 1
		}
	}
	if a+b == 2 {
		return false
	}
	return true
}

//------------------------Leetcode Problem 896------------------------
/*
 * https://leetcode.cn/problems/monotonic-array/
 * 执行用时：124ms 在所有Go提交中击败了57.87%的用户
 * 占用内存：8.7MB 在所有Go提交中击败了88.76%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", isMonotonic(nums))
	}
}
