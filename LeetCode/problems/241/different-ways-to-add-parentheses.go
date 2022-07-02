package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/different-ways-to-add-parentheses/
//------------------------Leetcode Problem 241------------------------
func diffWaysToCompute(expression string) (ans []int) {
	for i, c := range expression {
		if c == '+' || c == '-' || c == '*' {
			left := diffWaysToCompute(expression[:i])
			right := diffWaysToCompute(expression[i+1:])
			for _, l := range left {
				for _, r := range right {
					if c == '+' {
						ans = append(ans, l+r)
					} else if c == '-' {
						ans = append(ans, l-r)
					} else {
						ans = append(ans, l*r)
					}
				}
			}
		}
	}
	if len(ans) == 0 {
		n, _ := strconv.Atoi(expression)
		ans = append(ans, n)
	}
	return ans
}

//------------------------Leetcode Problem 241------------------------
/*
 * https://leetcode.cn/problems/different-ways-to-add-parentheses/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.1MB 在所有Go提交中击败了79.08%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		Printf("Output: %v\n", diffWaysToCompute(scanner.Text()))
	}
}
