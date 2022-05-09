package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func countBits(n int) []int {
	ans := make([]int, n+1)
	for i := 0; i <= n; i++ {
		ans[i] = ans[i>>1] + (i & 1)
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		ans := countBits(n)
		fmt.Printf("%v\n", ans)
	}
}


/*
 * 执行用时：4ms 在所有Go提交中击败了83.47%的用户
 * 占用内存：4.4MB 在所有Go提交中击败了100.00%的用户
**/
