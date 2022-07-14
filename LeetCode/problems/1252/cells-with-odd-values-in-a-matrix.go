package main

import (
	"bufio"
	. "fmt"
	"os"
)

//------------------------Leetcode Problem 1252------------------------
// https://leetcode.cn/problems/cells-with-odd-values-in-a-matrix/
func oddCells(m int, n int, indices [][]int) int {
	rows, cols := make([]int, m), make([]int, n)
	for _, p := range indices {
		rows[p[0]]++
		cols[p[1]]++
	}
	var oddx, oddy int
	for _, x := range rows {
		oddx += x & 1
	}
	for _, y := range cols {
		oddy += y & 1
	}
	return oddx*(n-oddy) + (m-oddx)*oddy
}

//------------------------Leetcode Problem 1252------------------------
/*
 * https://leetcode.cn/problems/cells-with-odd-values-in-a-matrix/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.3MB 在所有Go提交中击败了75.76%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var m, n int
		Sscanf(scanner.Text(), "%d %d", &m, &n)
		indices := make([][]int, m)
		for i := range indices {
			indices[i] = make([]int, n)
		}
		Printf("Output: %v\n", oddCells(m, n, indices))
	}
}
