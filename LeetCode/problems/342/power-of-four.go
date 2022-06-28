package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/power-of-four/
//------------------------Leetcode Problem 342------------------------
func isPowerOfFour(n int) bool {
	if n < 4 {
		return n == 1
	} else if n%4 == 0 {
		return isPowerOfFour(n / 4)
	} else {
		return false
	}
}

//------------------------Leetcode Problem 342------------------------
/*
 * https://leetcode.cn/problems/power-of-four/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.1MB 在所有Go提交中击败了71.20%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", isPowerOfFour(n))
	}
}
