package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/fibonacci-number/
//------------------------Leetcode Problem 509------------------------
func fib(n int) int {
	f, g := 0, 1
	for i := 0; i < n; i++ {
		g += f
		f = g - f
	}
	return f
}

//------------------------Leetcode Problem 509------------------------
/*
 * https://leetcode.cn/problems/fibonacci-number/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了25.61%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", fib(n))
	}
}
