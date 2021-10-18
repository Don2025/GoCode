package main

import (
	"bufio"
	"os"
	"strconv"
)

func findComplement(num int) int {
	var ans int
	for i := num; i > 0; i >>= 1 {
		ans = (ans << 1) + 1
	}
	return ans ^ num
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		println(findComplement(n))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.4MB 在所有Go提交中击败了89.90%的用户
**/
