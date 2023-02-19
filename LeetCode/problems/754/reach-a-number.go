package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

//------------------------Leetcode Problem 754------------------------
// https://leetcode.cn/problems/reach-a-number/
func reachNumber(target int) int {
	if target < 0 {
		target = -target
	}
	sum, step := 0, 0
	for sum < target || (sum-target)%2 != 0 {
		step++
		sum += step
	}
	return step
}

//------------------------Leetcode Problem 754------------------------
/*
 * https://leetcode.cn/problems/reach-a-number/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.8MB 在所有Go提交中击败了51.16%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		target, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %d\n", reachNumber(target))
	}
}
