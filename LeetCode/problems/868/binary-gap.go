package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/binary-gap/
//------------------------Leetcode Problem 868------------------------
func binaryGap(n int) int {
	var ans int
	for i, j := 0, -1; n > 0; i++ {
		if n&1 == 1 {
			if j != -1 {
				ans = max(ans, i-j)
			}
			j = i
		}
		n >>= 1
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//------------------------Leetcode Problem 868------------------------
/*
 * https://leetcode.cn/problems/binary-gap/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.8MB 在所有Go提交中击败了73.81%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %d\n", binaryGap(n))
	}
}
