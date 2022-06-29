package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"math"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/combination-sum-iv/
//------------------------Leetcode Problem 377------------------------
func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for _, x := range nums {
			if i >= x && dp[i-x] < math.MaxInt32 {
				dp[i] += dp[i-x]
			}
		}
	}
	return dp[target]
}

//------------------------Leetcode Problem 377------------------------
/*
 * https://leetcode.cn/problems/combination-sum-iv/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了51.26%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		scanner.Scan()
		target, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", combinationSum4(nums, target))
	}
}
