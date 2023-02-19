package main

import (
	"bufio"
	. "fmt"
	"os"
)

//------------------------Leetcode Problem 1106------------------------
// https://leetcode.cn/problems/parsing-a-boolean-expression/
func parseBoolExpr(expression string) bool {
	stack := []rune{}
	for _, c := range expression {
		if c == ',' {
			continue
		}
		if c != ')' {
			stack = append(stack, c)
			continue
		}
		var t, f int
		for stack[len(stack)-1] != '(' {
			val := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if val == 't' {
				t++
			} else {
				f++
			}
		}
		stack = stack[:len(stack)-1]
		op := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		c = 't'
		switch op {
		case '!':
			if f != 1 {
				c = 'f'
			}
			stack = append(stack, c)
		case '&':
			if f != 0 {
				c = 'f'
			}
			stack = append(stack, c)
		case '|':
			if t == 0 {
				c = 'f'
			}
			stack = append(stack, c)
		}
	}
	return stack[len(stack)-1] == 't'
}

//------------------------Leetcode Problem 1106------------------------
/*
 * https://leetcode.cn/problems/parsing-a-boolean-expression/
 * 执行用时：4ms 在所有Go提交中击败了7.14%的用户
 * 占用内存：2.3MB 在所有Go提交中击败了28.57%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		Printf("Output: %v\n", parseBoolExpr(scanner.Text()))
	}
}
