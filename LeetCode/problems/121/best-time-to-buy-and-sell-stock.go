package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"math"
	"os"
)

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/
//------------------------Leetcode Problem 121------------------------
func maxProfit(prices []int) int {
	lowPrice := math.MaxInt32
	var maxProfit int
	for _, price := range prices {
		lowPrice = min(lowPrice, price)
		maxProfit = max(maxProfit, price-lowPrice)
	}
	return maxProfit
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//------------------------Leetcode Problem 121------------------------
/*
 * https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/
 * 执行用时：124ms 在所有Go提交中击败了34.21%的用户
 * 占用内存：8.2MB 在所有Go提交中击败了30.56%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		prices := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", maxProfit(prices))
	}
}
