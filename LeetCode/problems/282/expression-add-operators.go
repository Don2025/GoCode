package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/expression-add-operators/
//------------------------Leetcode Problem 282------------------------
func addOperators(num string, target int) (ans []string) {
	n := len(num)
	if n == 0 {
		return
	}
	var backtrack func(expr []byte, i, res, mul int)
	backtrack = func(expr []byte, i, res, mul int) {
		if i == n {
			if res == target {
				ans = append(ans, string(expr))
			}
			return
		}
		signIndex := len(expr)
		if i > 0 {
			expr = append(expr, 0)
		}
		for j, val := i, 0; j < n && (j == i || num[i] != '0'); j++ {
			val = val*10 + int(num[j]-'0')
			expr = append(expr, num[j])
			if i == 0 {
				backtrack(expr, j+1, val, val)
			} else {
				expr[signIndex] = '+'
				backtrack(expr, j+1, res+val, val)
				expr[signIndex] = '-'
				backtrack(expr, j+1, res-val, -val)
				expr[signIndex] = '*'
				backtrack(expr, j+1, res-mul+mul*val, mul*val)
			}
		}
	}
	backtrack(make([]byte, 0, n<<1-1), 0, 0, 0)
	return
}

//------------------------Leetcode Problem 282------------------------
/*
 * https://leetcode.cn/problems/expression-add-operators/
 * 执行用时：8ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：3.1MB 在所有Go提交中击败了100.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		num := scanner.Text()
		scanner.Scan()
		target, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", addOperators(num, target))
	}
}
