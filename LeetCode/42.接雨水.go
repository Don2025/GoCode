package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func trap(height []int) int {
	n := len(height)
	left, right := make([]int, n), make([]int, n)
	for i := 1; i < n; i++ {
		left[i] = max(left[i-1], height[i-1])
	}
	for i := n - 2; i >= 0; i-- {
		right[i] = max(right[i+1], height[i+1])
	}
	var ans int
	for i := range height {
		level := min(left[i], right[i])
		ans += max(0, level-height[i])
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

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		height := stringArrayToIntArray(strings.Fields(input.Text()))
		println(trap(height))
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
 * 执行用时：8ms 在所有Go提交中击败了55.34%的用户
 * 占用内存：4.1MB 在所有Go提交中击败了31.65%的用户
**/
