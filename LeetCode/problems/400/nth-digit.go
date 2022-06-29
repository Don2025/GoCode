package main

import (
	"bufio"
	. "fmt"
	"math"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/nth-digit/
//------------------------Leetcode Problem 400------------------------
func findNthDigit(n int) int {
	d := 1
	for cnt := 9; n > d*cnt; cnt *= 10 {
		n -= d * cnt
		d++
	}
	index := n - 1
	start := int(math.Pow10(d - 1))
	num := start + index/d
	digitIndex := index % d
	return num / int(math.Pow10(d-digitIndex-1)) % 10
}

//------------------------Leetcode Problem 400------------------------
/*
 * https://leetcode.cn/problems/nth-digit/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了37.33%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", findNthDigit(n))
	}
}
