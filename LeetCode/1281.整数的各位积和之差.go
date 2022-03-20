package main

import (
	"bufio"
	"os"
	"strconv"
)

func subtractProductAndSum(n int) int {
	add, mul := 0, 1
	for i := n; i > 0; i /= 10 {
		t := i % 10
		add += t
		mul *= t
	}
	return mul - add
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		println(subtractProductAndSum(n))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.8MB 在所有Go提交中击败了58.95%的用户
**/
