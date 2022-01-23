package main

import (
	"bufio"
	"math"
	"os"
	"strings"
)

func average(salary []int) float64 {
	var sum, maxVal int
	minVal := math.MaxInt32
	for _, x := range salary {
		sum += x
		maxVal = max(maxVal, x)
		minVal = min(minVal, x)
	}
	ans := float64(sum-maxVal-minVal) / float64(len(salary)-2)
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

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		salary := stringArrayToIntArray(strings.Fields(input.Text()))
		println(average(salary))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了100.00%的用户
**/
