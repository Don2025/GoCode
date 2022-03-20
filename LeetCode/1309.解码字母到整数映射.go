package main

import (
	"bufio"
	"os"
	"strconv"
)

func freqAlphabets(s string) string {
	var ans string
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '#' {
			x, _ := strconv.Atoi(s[i-2 : i])
			ans = string('a'+x-1) + ans
			i -= 2
		} else {
			x, _ := strconv.Atoi(string(s[i]))
			ans = string('a'+x-1) + ans
		}
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		println(freqAlphabets(input.Text()))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.4MB 在所有Go提交中击败了28.00%的用户
**/
