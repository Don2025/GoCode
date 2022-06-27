package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/power-of-three/
//------------------------Leetcode Problem 326------------------------
func isPowerOfThree(n int) bool {
	if n < 3 {
		return n == 1
	}
	if n%3 == 0 {
		return isPowerOfThree(n / 3)
	}
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", isPowerOfThree(n))
	}
}

//------------------------Leetcode Problem 326------------------------
/*
 * https://leetcode.cn/problems/power-of-three/
 * 执行用时：20ms 在所有Go提交中击败了91.48%的用户
 * 占用内存：6.4MB 在所有Go提交中击败了70.33%的用户
**/
