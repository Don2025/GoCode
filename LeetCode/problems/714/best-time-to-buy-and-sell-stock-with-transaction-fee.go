package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"math"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/
//------------------------Leetcode Problem 714------------------------
func maxProfit(prices []int, fee int) int {
	var ans int
	minPrice := math.MaxInt32
	for i := 0; i < len(prices); i++ {
		minPrice = min(minPrice, prices[i])
		if prices[i] > minPrice+fee {
			ans += prices[i] - minPrice - fee
			minPrice = prices[i] - fee
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		prices := utils.StringToInts(scanner.Text())
		scanner.Scan()
		fee, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", maxProfit(prices, fee))
	}
}

//------------------------Leetcode Problem 714------------------------
/*
 * https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/
 * 执行用时：92ms 在所有Go提交中击败了33.87%的用户
 * 占用内存：7.6MB 在所有Go提交中击败了66.85%的用户
**/
