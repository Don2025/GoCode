package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/super-washing-machines/
//------------------------Leetcode Problem 517------------------------
func findMinMoves(machines []int) int {
	var total int
	for _, x := range machines {
		total += x
	}
	n := len(machines)
	if total%n != 0 {
		return -1
	}
	avg := total / n
	var ans, sum int
	for _, x := range machines {
		x -= avg
		sum += x
		ans = max(ans, max(abs(sum), x))
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(n int) int {
	if n < 0 {
		n = -n
	}
	return n
}

//------------------------Leetcode Problem 517------------------------
/*
 * https://leetcode.cn/problems/super-washing-machines/
 * 执行用时：4ms 在所有Go提交中击败了100.0%的用户
 * 占用内存：3.99MB 在所有Go提交中击败了83.31%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		machines := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", findMinMoves(machines))
	}
}
