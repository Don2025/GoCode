package main

import (
	"bufio"
	"os"
	"strconv"
)

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 1 / myPow(x, -n)
	}
	if n%2 == 1 {
		return x * myPow(x, n-1)
	}
	return myPow(x*x, n/2)
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		x, _ := strconv.ParseFloat(input.Text(), 32)
		input.Scan()
		n, _ := strconv.Atoi(input.Text())
		println(myPow(x, n))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2MB 在所有Go提交中击败了59.35%的用户
**/
