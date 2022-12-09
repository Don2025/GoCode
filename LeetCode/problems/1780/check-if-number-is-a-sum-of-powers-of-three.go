package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/check-if-number-is-a-sum-of-powers-of-three/
//------------------------Leetcode Problem 1780------------------------
func checkPowersOfThree(n int) bool {
	for n > 0 {
		if n%3 == 2 {
			return false
		}
		n /= 3
	}
	return true
}

//------------------------Leetcode Problem 1780------------------------
/*
 * https://leetcode.cn/problems/check-if-number-is-a-sum-of-powers-of-three/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.8MB 在所有Go提交中击败了100.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", checkPowersOfThree(n))
	}
}
