package main

import (
	"bufio"
	"os"
	"strconv"
)

/* 递归法直接TLE，超出时间限制啦
func numWays(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return numWays(n-1)+numWays(n-2)
}
*/

func numWays(n int) int {
	f, g := 1, 1
	for i := 0; i < n; i++ {
		f, g = g, (f+g)%1000000007
	}
	return f
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		println(numWays(n))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了69.09%的用户
**/
