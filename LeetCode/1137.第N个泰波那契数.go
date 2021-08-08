package main

import (
	"bufio"
	"os"
	"strconv"
)

/* Time Limit Exceeded
func tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	return tribonacci(n-3)+tribonacci(n-2)+tribonacci(n-1)
}
*/

func tribonacci(n int) int {
	t0, t1, t2 := 0, 1, 1
	for i := 0; i < n; i++ {
		t0, t1, t2 = t1, t2, t0+t1+t2
	}
	return t0
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		println(tribonacci(n))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了71.69%的用户
**/
