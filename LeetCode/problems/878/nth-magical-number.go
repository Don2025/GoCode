package main

import (
	"bufio"
	. "fmt"
	"os"
)

// https://leetcode.cn/problems/nth-magical-number/
//------------------------Leetcode Problem 878------------------------
func nthMagicalNumber(n int, a int, b int) int {
	const MOD = 1e9 + 7
	c := a / gcd(a, b) * b
	m := c/a + c/b - 1
	r := n % m
	res := c * (n / m) % MOD
	if r == 0 {
		return res
	}
	addA := a
	addB := b
	for i := 0; i < r-1; i++ {
		if addA < addB {
			addA += a
		} else {
			addB += b
		}
	}
	return int(res+min(addA, addB)%MOD) % MOD
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func gcd(a, b int) int {
	for a%b != 0 {
		a, b = b, a%b
	}
	return b
}

//------------------------Leetcode Problem 878------------------------
/*
 * https://leetcode.cn/problems/nth-magical-number/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.8MB 在所有Go提交中击败了100.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var n, a, b int
		Sscanf(scanner.Text(), "%d %d %d", &n, &a, &b)
		Printf("Output: %v\n", nthMagicalNumber(n, a, b))
	}
}
