package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/minimum-moves-to-equal-array-elements/
//------------------------Leetcode Problem 453------------------------
func minMoves(nums []int) int {
	var ans int
	min := nums[0]
	for _, num := range nums[1:] {
		if num < min {
			min = num
		}
	}
	for _, num := range nums {
		ans += num - min
	}
	return ans
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", minMoves(nums))
	}
}

//------------------------Leetcode Problem 453------------------------
/*
 * https://leetcode.cn/problems/minimum-moves-to-equal-array-elements/
 * 执行用时：36ms 在所有Go提交中击败了66.06%的用户
 * 占用内存：6.6MB 在所有Go提交中击败了98.17%的用户
**/
