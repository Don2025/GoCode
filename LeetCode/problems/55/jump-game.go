package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/jump-game/
//------------------------Leetcode Problem 55------------------------
func canJump(nums []int) bool {
	var maxDis int
	for i := range nums {
		if i <= maxDis {
			maxDis = max(maxDis, i+nums[i])
			if maxDis >= len(nums)-1 {
				return true
			}
		}
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//------------------------Leetcode Problem 55------------------------
/*
 * https://leetcode.cn/problems/jump-game/
 * 执行用时：60ms 在所有Go提交中击败了54.43%的用户
 * 占用内存：6.9MB 在所有Go提交中击败了64.35%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", canJump(nums))
	}
}
