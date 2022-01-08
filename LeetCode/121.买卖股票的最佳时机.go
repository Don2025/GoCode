package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

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

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		prices := stringArrayToIntArray(strings.Fields(input.Text()))
		println(maxProfit(prices))
	}
}

func stringArrayToIntArray(strings []string) []int {
	var arr []int
	for _, x := range strings {
		n, _ := strconv.Atoi(x)
		arr = append(arr, n)
	}
	return arr
}

/*
 * 执行用时：124ms 在所有Go提交中击败了34.21%的用户
 * 占用内存：8.2MB 在所有Go提交中击败了30.56%的用户
**/
