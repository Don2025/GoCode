package main

import (
	"bufio"
	"os"
	"strconv"
)

func minSteps(n int) int {
	var ans int
	for i := 2; i*i <= n; i++ {
		for n%i == 0 {
			n /= i
			ans += i
		}
	}
	if n > 1 {
		ans += n
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		println(minSteps(n))
	}
}

/*
 * 执行用时：0-ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了90.24%的用户
**/
