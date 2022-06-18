package main

import (
	"bufio"
	"github.com/Don2025/GoCode/utils"
	"os"
)

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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		prices := utils.StringToInts(scanner.Text())
		println(maxProfit(prices))
	}
}

/*
 * 执行用时：4ms 在所有Go提交中击败了91.51%的用户
 * 占用内存：3MB 在所有Go提交中击败了99.95%的用户
**/
