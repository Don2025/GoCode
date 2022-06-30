package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/prime-arrangements/
//------------------------Leetcode Problem 1175------------------------

const MOD = 1e9 + 7

func numPrimeArrangements(n int) int {
	var cnt int
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			cnt++
		}
	}
	return factorial(cnt) * factorial(n-cnt) % MOD
}

func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func factorial(n int) int {
	f := 1
	for i := 1; i <= n; i++ {
		f = f * i % MOD
	}
	return f
}

//------------------------Leetcode Problem 1175------------------------
/*
 * https://leetcode.cn/problems/prime-arrangements/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了8.33%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", numPrimeArrangements(n))
	}
}
