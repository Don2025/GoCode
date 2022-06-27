package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/number-of-1-bits/
//------------------------Leetcode Problem 191------------------------
func hammingWeight(num uint32) int {
	cnt := 0
	for num != 0 {
		num &= (num - 1)
		cnt++
	}
	return cnt
}

//------------------------Leetcode Problem 191------------------------
/*
 * https://leetcode.cn/problems/number-of-1-bits/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了100.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		num, _ := strconv.ParseUint(scanner.Text(), 2, 32)
		Printf("Output: %v\n", hammingWeight(uint32(num)))
	}
}
