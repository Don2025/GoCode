package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/number-complement/
//------------------------Leetcode Problem 476------------------------

func findComplement(num int) int {
	var ans int
	for i := num; i > 0; i >>= 1 {
		ans = (ans << 1) + 1
	}
	return ans ^ num
}

//------------------------Leetcode Problem 476------------------------
/*
 * https://leetcode.cn/problems/number-complement/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.4MB 在所有Go提交中击败了89.90%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", findComplement(n))
	}
}
