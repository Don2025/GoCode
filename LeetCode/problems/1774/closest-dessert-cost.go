package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/closest-dessert-cost/
//------------------------Leetcode Problem 1774------------------------
func closestCost(baseCosts []int, toppingCosts []int, target int) int {
	ans := baseCosts[0]
	for _, c := range baseCosts {
		ans = min(ans, c)
	}
	var dfs func(int, int)
	dfs = func(i, curCost int) {
		if abs(ans-target) < curCost-target {
			return
		} else if abs(ans-target) >= abs(curCost-target) {
			if abs(ans-target) > abs(curCost-target) {
				ans = curCost
			} else {
				ans = min(ans, curCost)
			}
		}
		if i == len(toppingCosts) {
			return
		}
		dfs(i+1, curCost)
		dfs(i+1, curCost+toppingCosts[i])
		dfs(i+1, curCost+toppingCosts[i]*2)
	}
	for _, c := range baseCosts {
		dfs(0, c)
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(n int) int {
	if n < 0 {
		n = -n
	}
	return n
}

//------------------------Leetcode Problem 1774------------------------
/*
 * https://leetcode.cn/problems/closest-dessert-cost/
 * 执行用时：4ms 在所有Go提交中击败了82.35%的用户
 * 占用内存：1.8MB 在所有Go提交中击败了94.12%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		baseCosts := utils.StringToInts(scanner.Text())
		scanner.Scan()
		toppingCosts := utils.StringToInts(scanner.Text())
		scanner.Scan()
		target, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", closestCost(baseCosts, toppingCosts, target))
	}
}
