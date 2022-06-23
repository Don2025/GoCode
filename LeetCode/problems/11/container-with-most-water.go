package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/container-with-most-water/
//------------------------Leetcode Problem 11------------------------
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

//------------------------Leetcode Problem 11------------------------
/*
 * https://leetcode.cn/problems/container-with-most-water/
 * 执行用时：68ms 在所有Go提交中击败了78.07%的用户
 * 占用内存：8.3MB 在所有Go提交中击败了71.99%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		height := utils.StringToInts(scanner.Text())
		Printf("%v\n", maxArea(height))
	}
}
