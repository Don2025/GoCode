package main

import (
	"bufio"
	. "fmt"
	"math"
	"os"
)

// https://leetcode.cn/problems/omKAoA/
// ------------------------剑指 Offer II Problem 94------------------------
func minCut(s string) int {
	n := len(s)
	g := make([][]bool, n)
	for i := range g {
		g[i] = make([]bool, n)
		for j := range g[i] {
			g[i][j] = true
		}
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			g[i][j] = s[i] == s[j] && g[i+1][j-1]
		}
	}

	f := make([]int, n)
	for i := range f {
		if g[0][i] {
			continue
		}
		f[i] = math.MaxInt64
		for j := 0; j < i; j++ {
			if g[j+1][i] && f[j]+1 < f[i] {
				f[i] = f[j] + 1
			}
		}
	}
	return f[n-1]
}

// ------------------------剑指 Offer II Problem 94------------------------
/*
 * https://leetcode.cn/problems/omKAoA/
 * 执行用时：12ms 在所有Go提交中击败了83.64%的用户
 * 占用内存：7.3MB 在所有Go提交中击败了62.73%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	Printf("Input a line of string:")
	for scanner.Scan() {
		Printf("Output: %v\n", minCut(scanner.Text()))
		Printf("Input a line of string:")
	}
}
