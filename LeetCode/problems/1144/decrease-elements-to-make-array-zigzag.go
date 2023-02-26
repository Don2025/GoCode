package main

import (
	"bufio"
	"fmt"
	"github.com/Don2025/GoCode/utils"
	"math"
	"os"
)

// https://leetcode.cn/problems/decrease-elements-to-make-array-zigzag/
//------------------------Leetcode Problem 1144------------------------
func movesToMakeZigzag(nums []int) int {
	ans := make([]int, 2)
	for i := range nums {
		l, r := math.MaxInt, math.MaxInt
		if i > 0 {
			l = nums[i-1]
		}
		if i < len(nums)-1 {
			r = nums[i+1]
		}
		ans[i%2] += max(0, nums[i]-min(l, r)+1)
	}
	return min(ans[0], ans[1])
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

//------------------------Leetcode Problem 1144------------------------
/*
 * https://leetcode.cn/problems/decrease-elements-to-make-array-zigzag/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2MB 在所有Go提交中击败了12.50%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		fmt.Printf("%d\n", movesToMakeZigzag(nums))
	}
}
