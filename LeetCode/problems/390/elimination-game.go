package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/elimination-game/
//------------------------Leetcode Problem 390------------------------
func lastRemaining(n int) int {
	ans, step := 1, 1
	flag := true // true左false右
	for n > 1 {
		if flag || (n&1) == 1 {
			ans += step
		}
		flag = !flag
		step <<= 1
		n >>= 1
	}
	return ans
}

//------------------------Leetcode Problem 390------------------------
/*
 * https://leetcode.cn/problems/elimination-game/
 * 执行用时：12ms 在所有Go提交中击败了12.82%的用户
 * 占用内存：2.7MB 在所有Go提交中击败了100.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", lastRemaining(n))
	}
}
