package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/iIQa4I/
// ------------------------剑指 Offer II Problem 38------------------------
func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	ans := make([]int, n)
	var stack []int
	for i := 0; i < n; i++ {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			preIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[preIndex] = i - preIndex
		}
		stack = append(stack, i)
	}
	return ans
}

// ------------------------剑指 Offer II Problem 38------------------------
/*
 * https://leetcode.cn/problems/iIQa4I/
 * 执行用时：108ms 在所有Go提交中击败了88.00%的用户
 * 占用内存：8.6MB 在所有Go提交中击败了84.80%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		temperatures := utils.StringToInts(scanner.Text())
		Printf("%v\n", dailyTemperatures(temperatures))
	}
}
