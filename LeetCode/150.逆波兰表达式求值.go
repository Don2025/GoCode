package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func evalRPN(tokens []string) int {
	var stack []int
	for _, token := range tokens {
		val, err := strconv.Atoi(token)
		if err == nil {
			stack = append(stack, val)
		} else {
			num1, num2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			switch token {
			case "+":
				stack = append(stack, num1+num2)
			case "-":
				stack = append(stack, num1-num2)
			case "*":
				stack = append(stack, num1*num2)
			default:
				stack = append(stack, num1/num2)
			}
		}
	}
	return stack[0]
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		tokens := strings.Fields(input.Text())
		println(evalRPN(tokens))
	}
}

/*
 * 执行用时：4ms 在所有Go提交中击败了83.86%的用户
 * 占用内存：4.6MB 在所有Go提交中击败了29.33%的用户
**/
