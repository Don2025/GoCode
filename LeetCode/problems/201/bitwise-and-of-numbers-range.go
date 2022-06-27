package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
	"strings"
)

// https://leetcode.cn/problems/bitwise-and-of-numbers-range/
//------------------------Leetcode Problem 201------------------------
func rangeBitwiseAnd(left int, right int) int {
	var shift int
	for left < right {
		left >>= 1
		right >>= 1
		shift++
	}
	return left << shift
}

//------------------------Leetcode Problem 201------------------------
/*
 * https://leetcode.cn/problems/bitwise-and-of-numbers-range/
 * 执行用时：4ms 在所有Go提交中击败了97.94%的用户
 * 占用内存：5MB 在所有Go提交中击败了63.79%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		arr := strings.Fields(scanner.Text())
		left, _ := strconv.Atoi(arr[0])
		right, _ := strconv.Atoi(arr[1])
		Printf("Output: %v\n", rangeBitwiseAnd(left, right))
	}
}
