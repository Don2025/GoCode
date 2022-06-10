package main

import (
	"bufio"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strings"
)

func maxArea(height []int) int {
	n := len(height)
	if n < 1 {
		return -1
	}
	i, j := 0, n-1
	var ans int
	for i < j {
		h := min(height[i], height[j])
		ans = max(ans, h*(j-i))
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
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
		height := utils.StringArrayToIntArray(strings.Fields(input.Text()))
		println(maxArea(height))
	}
}

/*
 * 执行用时：68ms 在所有Go提交中击败了78.07%的用户
 * 占用内存：8.3MB 在所有Go提交中击败了71.99%的用户
**/
