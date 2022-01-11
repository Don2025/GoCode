package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

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
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		prices := stringArrayToIntArray(strings.Fields(input.Text()))
		input.Scan()
		fee, _ := strconv.Atoi(input.Text())
		println(maxProfit(prices, fee))
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
 * 执行用时：92ms 在所有Go提交中击败了33.87%的用户
 * 占用内存：7.6MB 在所有Go提交中击败了66.85%的用户
**/
