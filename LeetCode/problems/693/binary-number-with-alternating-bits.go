package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/binary-number-with-alternating-bits/
//------------------------Leetcode Problem 693------------------------
func hasAlternatingBits(n int) bool {
	n = n ^ (n >> 1)
	return (n & (n + 1)) == 0
}

//------------------------Leetcode Problem 693------------------------
/*
 * https://leetcode.cn/problems/binary-number-with-alternating-bits/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.8MB 在所有Go提交中击败了85.33%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", hasAlternatingBits(n))
	}
}
