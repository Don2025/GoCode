package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

func maxScoreSightseeingPair(values []int) int {
	left, ans := values[0], math.MinInt32
	for i := 1; i < len(values); i++ {
		ans = max(ans, left+values[i]-i)
		left = max(left, values[i]+i)
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
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		values := stringArrayToIntArray(strings.Fields(input.Text()))
		println(maxScoreSightseeingPair(values))
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
 * 执行用时：40ms 在所有Go提交中击败了99.16%的用户
 * 占用内存：7.2MB 在所有Go提交中击败了74.26%的用户
**/
