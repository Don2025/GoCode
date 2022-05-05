package main

import (
	"bufio"
	"os"
)

func titleToNumber(columnTitle string) int {
	var ans int
	for _, x := range columnTitle {
		ans *= 26
		ans += int(x - 'A' + 1)
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		println(titleToNumber(input.Text()))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2MB 在所有Go提交中击败了99.75%的用户
**/
