package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/count-number-of-maximum-bitwise-or-subsets/
//------------------------Leetcode Problem 2044------------------------
func countMaxOrSubsets(nums []int) int {
	var ans, maxOr int
	var dfs func(int, int)
	dfs = func(pos, or int) {
		if len(nums) == pos {
			if or > maxOr {
				maxOr = or
				ans = 1
			} else if or == maxOr {
				ans++
			}
			return
		}
		dfs(pos+1, or|nums[pos])
		dfs(pos+1, or)
	}
	dfs(0, 0)
	return ans
}

//------------------------Leetcode Problem 2044------------------------
/*
 * https://leetcode.cn/problems/count-number-of-maximum-bitwise-or-subsets/
 * 执行用时：8ms 在所有Go提交中击败了65.79%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了81.58%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		Printf("Output: %d\n", countMaxOrSubsets(nums))
	}
}
