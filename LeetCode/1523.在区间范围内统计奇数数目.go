package main

import (
	"bufio"
	"os"
	"strconv"
)

func countOdds(low int, high int) int {
	var f func(int) int
	f = func(i int) int {
		return (i + 1) >> 1
	}
	return f(high) - f(low-1)
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		low, _ := strconv.Atoi(input.Text())
		input.Scan()
		high, _ := strconv.Atoi(input.Text())
		println(countOdds(low, high))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了18.48%的用户
**/
