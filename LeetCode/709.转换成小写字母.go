package main

import (
	"bufio"
	"os"
)

func toLowerCase(s string) string {
	ans := []byte(s)
	for i := range ans {
		if ans[i] >= 'A' && ans[i] <= 'Z' {
			ans[i] |= 32
		}
	}
	return string(ans)
}

/* 大写变小写 小写变大写 c ^= 32
 * 大小写都变小写 c |= 32
 * 大小写都变大写 c &= -33
 */

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		println(toLowerCase(input.Text()))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2MB 在所有Go提交中击败了100.00%的用户
**/
