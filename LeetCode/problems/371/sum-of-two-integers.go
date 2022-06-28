package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
	"strings"
)

// https://leetcode.cn/problems/sum-of-two-integers/
//------------------------Leetcode Problem 371------------------------
func getSum(a int, b int) int {
	for b != 0 {
		carry := uint(a&b) << 1
		a ^= b
		b = int(carry)
	}
	return a
}

//------------------------Leetcode Problem 371------------------------
/*
 * https://leetcode.cn/problems/sum-of-two-integers/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.8MB 在所有Go提交中击败了43.56%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		arr := strings.Fields(scanner.Text())
		a, _ := strconv.Atoi(arr[0])
		b, _ := strconv.Atoi(arr[1])
		Printf("Output: %v\n", getSum(a, b))
	}
}
