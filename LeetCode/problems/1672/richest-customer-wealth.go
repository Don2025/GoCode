package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/richest-customer-wealth/
//------------------------Leetcode Problem 1672------------------------
func maximumWealth(accounts [][]int) int {
	var richest int
	for _, account := range accounts {
		var wealth int
		for _, x := range account {
			wealth += x
		}
		richest = max(richest, wealth)
	}
	return richest
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//------------------------Leetcode Problem 1672------------------------
/*
 * https://leetcode.cn/problems/richest-customer-wealth/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.9MB 在所有Go提交中击败了76.67%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		accounts := make([][]int, n)
		for i := 0; i < n; i++ {
			scanner.Scan()
			accounts[i] = utils.StringToInts(scanner.Text())
		}
		Printf("Output: %v\n", maximumWealth(accounts))
	}
}
