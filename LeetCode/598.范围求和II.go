package main

import (
	"bufio"
	"math"
	"os"
)

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

/*
 * 执行用时：4ms 在所有Go提交中击败了95.45%的用户
 * 占用内存：3.7MB 在所有Go提交中击败了100.00%的用户
**/
