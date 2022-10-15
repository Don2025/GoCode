package main

import (
	"bufio"
	. "fmt"
	"os"
)

// https://leetcode.cn/problems/distinct-subsequences-ii/
//------------------------Leetcode Problem 940------------------------
func distinctSubseqII(s string) (ans int) {
	const mod int = 1e9 + 7
	g := make([]int, 26)
	for _, c := range s {
		total := 1
		for _, v := range g {
			total = (total + v) % mod
		}
		g[c-'a'] = total
	}
	for _, v := range g {
		ans = (ans + v) % mod
	}
	return ans
}

//------------------------Leetcode Problem 940------------------------
/*
 * https://leetcode.cn/problems/distinct-subsequences-ii/
 * 执行用时：4ms 在所有Go提交中击败了69.48%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了78.06%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		Printf("Output: %d\n", distinctSubseqII(scanner.Text()))
	}
}
