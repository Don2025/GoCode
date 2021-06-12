package main

import (
	"bufio"
	"os"
	"strconv"
)

func sumNums(n int) int {
	if n == 0 {
		return 0
	}
	return n + sumNums(n-1)
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		println(sumNums(n))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.7MB 在所有Go提交中击败了55.60%的用户
**/
