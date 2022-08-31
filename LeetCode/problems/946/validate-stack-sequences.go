package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/validate-stack-sequences/
// ------------------------Leetcode Problem 946------------------------
func validateStackSequences(pushed []int, popped []int) bool {
	var stack []int
	for _, x := range pushed {
		stack = append(stack, x)
		for len(stack) > 0 && stack[len(stack)-1] == popped[0] {
			stack = stack[:len(stack)-1]
			popped = popped[1:]
		}
	}
	return len(stack) == 0
}

// ------------------------Leetcode Problem 946------------------------
/*
 * https://leetcode.cn/problems/validate-stack-sequences/
 * 执行用时：4ms 在所有Go提交中击败了95.17%的用户
 * 占用内存：3.6MB 在所有Go提交中击败了52.17%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		pushed := utils.StringToInts(scanner.Text())
		scanner.Scan()
		popped := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", validateStackSequences(pushed, popped))
	}
}
