package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

func maxChunksToSorted(arr []int) int {
	var stack []int
	for _, x := range arr {
		if len(stack) == 0 || x >= stack[len(stack)-1] {
			stack = append(stack, x)
		} else {
			mx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			for len(stack) > 0 && stack[len(stack)-1] > x {
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, mx)
		}
	}
	return len(stack)
}

//------------------------Leetcode Problem 768------------------------
/*
 * https://leetcode.cn/problems/min-cost-climbing-stairs/
 * 执行用时：4ms 在所有Go提交中击败了96.15%的用户
 * 占用内存：3.6MB 在所有Go提交中击败了80.77%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		arr := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", maxChunksToSorted(arr))
	}
}
