package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/coin-change-2/
//------------------------Leetcode Problem 518------------------------
func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := 1; i <= amount; i++ {
			if i >= coin {
				dp[i] = dp[i] + dp[i-coin]
			}
		}
	}
	return dp[amount]
}

//------------------------Leetcode Problem 518------------------------
/*
 * https://leetcode.cn/problems/coin-change-2/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.3MB 在所有Go提交中击败了62.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		coins := utils.StringToInts(scanner.Text())
		scanner.Scan()
		amount, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", change(amount, coins))
	}
}
