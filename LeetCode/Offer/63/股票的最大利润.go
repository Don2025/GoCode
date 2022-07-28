package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/gu-piao-de-zui-da-li-run-lcof/
// ------------------------剑指 Offer I Problem 63------------------------
func maxProfit(prices []int) int {
	minprice := 0x3f3f3f
	var ans int
	for _, price := range prices {
		ans = max(ans, price-minprice)
		minprice = min(minprice, price)
	}
	return ans
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

// ------------------------剑指 Offer I Problem 63------------------------
/*
 * https://leetcode.cn/problems/gu-piao-de-zui-da-li-run-lcof/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.9MB 在所有Go提交中击败了100.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		prices := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", maxProfit(prices))
	}
}
