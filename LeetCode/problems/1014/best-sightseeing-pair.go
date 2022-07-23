package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"math"
	"os"
)

// https://leetcode.cn/problems/best-sightseeing-pair/
//------------------------Leetcode Problem 1014------------------------
func maxScoreSightseeingPair(values []int) int {
	left, ans := values[0], math.MinInt32
	for i := 1; i < len(values); i++ {
		ans = max(ans, left+values[i]-i)
		left = max(left, values[i]+i)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//------------------------Leetcode Problem 1014------------------------
/*
 * https://leetcode.cn/problems/best-sightseeing-pair/
 * 执行用时：40ms 在所有Go提交中击败了99.16%的用户
 * 占用内存：7.2MB 在所有Go提交中击败了74.26%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		values := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", maxScoreSightseeingPair(values))
	}
}
