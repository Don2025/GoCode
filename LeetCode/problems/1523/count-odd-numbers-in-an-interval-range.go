package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/count-odd-numbers-in-an-interval-range/
//------------------------Leetcode Problem 1523------------------------
func countOdds(low int, high int) int {
	var f func(int) int
	f = func(i int) int {
		return (i + 1) >> 1
	}
	return f(high) - f(low-1)
}

//------------------------Leetcode Problem 1523------------------------
/*
 * https://leetcode.cn/problems/count-odd-numbers-in-an-interval-range/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了18.48%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		low, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		high, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", countOdds(low, high))
	}
}
