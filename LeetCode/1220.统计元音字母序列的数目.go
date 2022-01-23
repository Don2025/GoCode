package main

import (
	"bufio"
	"os"
	"strconv"
)

func countVowelPermutation(n int) int {
	const mod = 1e9 + 7
	a, e, i, o, u := 1, 1, 1, 1, 1
	for j := 1; j < n; j++ {
		a, e, i, o, u = e%mod, (a+i)%mod, (a+e+o+u)%mod, (i+u)%mod, a
	}
	return (a + e + i + o + u) % mod
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		println(countVowelPermutation(n))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了100.00%的用户
**/
