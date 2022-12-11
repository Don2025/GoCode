package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/minimum-operations-to-make-the-array-increasing/
//------------------------Leetcode Problem 1827------------------------
func minOperations(nums []int) int {
	ans := 0
	pre := nums[0] - 1
	for _, num := range nums {
		pre = max(pre+1, num)
		ans += pre - num
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//------------------------Leetcode Problem 1827------------------------
/*
 * https://leetcode.cn/problems/minimum-operations-to-make-the-array-increasing/
 * 执行用时：12ms 在所有Go提交中击败了59.52%的用户
 * 占用内存：5.3MB 在所有Go提交中击败了100.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", minOperations(nums))
	}
}
