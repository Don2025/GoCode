package main

import (
	"bufio"
	. "fmt"
	"math"
	"os"
)

// https://leetcode.cn/problems/range-addition-ii/
//------------------------Leetcode Problem 598------------------------
func maxCount(m int, n int, ops [][]int) int {
	if len(ops) == 0 {
		return m * n
	}
	a, b := math.MaxInt64, math.MaxInt64
	for i := 0; i < len(ops); i++ {
		a, b = min(a, ops[i][0]), min(b, ops[i][1])
	}
	return a * b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//------------------------Leetcode Problem 598------------------------
/*
 * https://leetcode.cn/problems/range-addition-ii/
 * 执行用时：4ms 在所有Go提交中击败了95.45%的用户
 * 占用内存：3.7MB 在所有Go提交中击败了100.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var m, n int
		var ops [][]int
		_, _ = Sscanf(scanner.Text(), "%d %d", &m, &n)
		for scanner.Scan() {
			var a, b int
			_, err := Sscanf(scanner.Text(), "%d %d", &a, &b)
			if err != nil {
				break
			}
			ops = append(ops, []int{a, b})
		}
		Printf("Output: %v\n", maxCount(m, n, ops))
	}
}
