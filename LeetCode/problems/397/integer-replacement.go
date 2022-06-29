package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/integer-replacement/
//------------------------Leetcode Problem 397------------------------
func integerReplacement(n int) int {
	if n == 1 {
		return 0
	}
	if n&1 == 0 {
		return 1 + integerReplacement(n/2)
	} else {
		return 1 + min(integerReplacement(n+1), integerReplacement(n-1))
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//------------------------Leetcode Problem 397------------------------
/*
 * https://leetcode.cn/problems/integer-replacement/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了100.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", integerReplacement(n))
	}
}
