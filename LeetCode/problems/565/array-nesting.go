package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/array-nesting/
//------------------------Leetcode Problem 589------------------------
func arrayNesting(nums []int) (ans int) {
	for i := range nums {
		var cnt int
		for nums[i] != -1 {
			i, nums[i] = nums[i], -1
			cnt++
		}
		ans = max(ans, cnt)
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//------------------------Leetcode Problem 565------------------------
/*
 * https://leetcode.cn/problems/array-nesting/
 * 执行用时：100ms 在所有Go提交中击败了83.89%的用户
 * 占用内存：9MB 在所有Go提交中击败了54.23%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", arrayNesting(nums))
	}
}
