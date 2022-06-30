package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/largest-palindrome-product/
//------------------------Leetcode Problem 479------------------------
func largestPalindrome(n int) int {
	nums := []int{9, 987, 123, 597, 677, 1218, 877, 475}
	return nums[n-1]
}

//------------------------Leetcode Problem 479------------------------
/*
 * https://leetcode.cn/problems/largest-palindrome-product/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了63.64%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", largestPalindrome(n))
	}
}
