package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/
//------------------------Leetcode Problem 122------------------------
func maxProfit(prices []int) int {
	var ans int
	for i := 1; i < len(prices); i++ {
		ans += max(0, prices[i]-prices[i-1])
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//------------------------Leetcode Problem 122------------------------
/*
 * https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/
 * 执行用时：4ms 在所有Go提交中击败了91.51%的用户
 * 占用内存：3MB 在所有Go提交中击败了99.95%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		prices := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", maxProfit(prices))
	}
}
