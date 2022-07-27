package main

import (
	"bufio"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/count-vowels-permutation/
//------------------------Leetcode Problem 1220------------------------
func countVowelPermutation(n int) int {
	const mod = 1e9 + 7
	a, e, i, o, u := 1, 1, 1, 1, 1
	for j := 1; j < n; j++ {
		a, e, i, o, u = e%mod, (a+i)%mod, (a+e+o+u)%mod, (i+u)%mod, a
	}
	return (a + e + i + o + u) % mod
}

//------------------------Leetcode Problem 1220------------------------
/*
 * https://leetcode.cn/problems/count-vowels-permutation/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了100.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		println(countVowelPermutation(n))
	}
}
