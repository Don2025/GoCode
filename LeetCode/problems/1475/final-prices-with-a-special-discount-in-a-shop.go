package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/final-prices-with-a-special-discount-in-a-shop/
//------------------------Leetcode Problem 1475------------------------
func finalPrices(prices []int) []int {
	ans := make([]int, len(prices))
	var stack []int
	for i := range prices {
		for len(stack) > 0 && prices[i] <= prices[stack[len(stack)-1]] {
			ans[stack[len(stack)-1]] = prices[stack[len(stack)-1]] - prices[i]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	for len(stack) > 0 {
		ans[stack[len(stack)-1]] = prices[stack[len(stack)-1]]
		stack = stack[:len(stack)-1]
	}
	return ans
}

//------------------------Leetcode Problem 1475------------------------
/*
 * https://leetcode.cn/problems/final-prices-with-a-special-discount-in-a-shop/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.9MB 在所有Go提交中击败了34.18%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		prices := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", finalPrices(prices))
	}
}
