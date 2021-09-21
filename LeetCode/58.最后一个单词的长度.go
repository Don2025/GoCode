package main

import (
	"bufio"
	"os"
	"strings"
)

func lengthOfLastWord(s string) int {
	arr := strings.Fields(s)
	return len(arr[len(arr)-1])
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		println(lengthOfLastWord(input.Text()))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.1MB 在所有Go提交中击败了28.43%的用户
**/
