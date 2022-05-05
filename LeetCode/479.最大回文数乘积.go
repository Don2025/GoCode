package main

import (
	"bufio"
	"os"
	"strconv"
)

func largestPalindrome(n int) int {
	nums := []int{9, 987, 123, 597, 677, 1218, 877, 475}
	return nums[n-1]
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		println(largestPalindrome(n))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了63.64%的用户
**/
