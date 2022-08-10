package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
	"unicode"
)

// https://leetcode.cn/problems/solve-the-equation/
//------------------------Leetcode Problem 640------------------------
func solveEquation(equation string) string {
	factor, val := 0, 0
	i, n, sign := 0, len(equation), 1 // 等式左边默认系数为正
	for i < n {
		if equation[i] == '=' {
			sign = -1 // 等式右边默认系数为负
			i++
			continue
		}
		s := sign
		if equation[i] == '+' { // 去掉前面的符号
			i++
		} else if equation[i] == '-' {
			s = -s
			i++
		}
		num, valid := 0, false
		for i < n && unicode.IsDigit(rune(equation[i])) {
			valid = true
			num = num*10 + int(equation[i]-'0')
			i++
		}
		if i < n && equation[i] == 'x' { // 变量
			if valid {
				s *= num
			}
			factor += s
			i++
		} else { // 数值
			val += s * num
		}
	}
	if factor == 0 {
		if val == 0 {
			return "Infinite solutions"
		}
		return "No solution"
	}
	return "x=" + strconv.Itoa(-val/factor)
}

//------------------------Leetcode Problem 640------------------------
/*
 * https://leetcode.cn/problems/solve-the-equation/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.8MB 在所有Go提交中击败了100.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		equation := scanner.Text()
		Printf("Output: %v\n", solveEquation(equation))
	}
}
