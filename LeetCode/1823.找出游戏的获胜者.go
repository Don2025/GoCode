package main

import (
	"bufio"
	"os"
	"strconv"
)

func findTheWinner(n int, k int) int {
	var p int
	for i := 2; i <= n; i++ {
		p = (p+k)%i
	}
	return p+1s
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		input.Scan()
		k, _ := strconv.Atoi(input.Text())
		println(findTheWinner(n, k))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了79.19%的用户
**/
