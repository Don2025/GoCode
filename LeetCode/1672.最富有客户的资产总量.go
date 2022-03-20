package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

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

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		accounts := make([][]int, n)
		for i := 0; i < n; i++ {
			input.Scan()
			accounts[i] = stringArrayToIntArray(strings.Fields(input.Text()))
		}
		println(maximumWealth(accounts))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.9MB 在所有Go提交中击败了76.67%的用户
**/
