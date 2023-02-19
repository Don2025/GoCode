package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/partition-array-into-disjoint-intervals/
//------------------------Leetcode Problem 915------------------------

func partitionDisjoint(nums []int) int {
	max := nums[0]
	leftMax := nums[0]
	ans := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] < leftMax {
			ans = i + 1
			leftMax = max
		} else if nums[i] > max {
			max = nums[i]
		}
	}
	return ans
}

//------------------------Leetcode Problem 915------------------------
/*
 * https://leetcode.cn/problems/partition-array-into-disjoint-intervals/
 * 执行用时：144ms 在所有Go提交中击败了11.69%的用户
 * 占用内存：8.1MB 在所有Go提交中击败了98.70%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		Printf("Output: %d\n", partitionDisjoint(nums))
	}
}
