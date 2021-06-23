package main

import (
	"bufio"
	"os"
	"strconv"
)

func convertToTitle(columnNumber int) string {
	var ans string
	if columnNumber <= 0 {
		return ans
	}
	for columnNumber > 0 {
		columnNumber--
		ans = string(columnNumber%26+'A') + ans
		columnNumber /= 26
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		println(convertToTitle(n))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了99.48%的用户
**/
