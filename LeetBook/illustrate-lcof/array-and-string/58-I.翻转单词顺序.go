package main

import (
	"bufio"
	"os"
	"strings"
)

func reverseWords(s string) string {
	var ans string
	if len(s) == 0 {
		return ans
	}
	arr := strings.Fields(s)
	for _, x := range arr {
		ans = x + " " + ans
	}
	return strings.TrimSpace(ans)
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		println(reverseWords(input.Text()))
	}
}

/*
 * 执行用时：8ms 在所有Go提交中击败了28.62%的用户
 * 占用内存：7.3MB 在所有Go提交中击败了12.96%的用户
**/
