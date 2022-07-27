package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/running-sum-of-1d-array/
//------------------------Leetcode Problem 1480------------------------
func runningSum(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	return nums
}

//------------------------Leetcode Problem 1480------------------------
/*
 * https://leetcode.cn/problems/running-sum-of-1d-array/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.7MB 在所有Go提交中击败了100.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		ans := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", runningSum(ans))
	}
}
