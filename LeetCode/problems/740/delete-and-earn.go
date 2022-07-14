package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/delete-and-earn/
//------------------------Leetcode Problem 740------------------------
func deleteAndEarn(nums []int) int {
	cnt := make(map[int]int)
	var maxNum int
	for _, num := range nums {
		cnt[num] += num
		maxNum = max(maxNum, num)
	}
	dp := make([]int, maxNum+1)
	dp[0], dp[1] = 0, cnt[1]
	for i := 2; i <= maxNum; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+cnt[i])
	}
	return dp[maxNum]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//------------------------Leetcode Problem 740------------------------
/*
 * https://leetcode.cn/problems/delete-and-earn/
 * 执行用时：4ms 在所有Go提交中击败了83.58%的用户
 * 占用内存：3.7MB 在所有Go提交中击败了22.55%的用户
**/

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		nums := utils.StringToInts(input.Text())
		Printf("Output: %v\n", deleteAndEarn(nums))
	}
}
