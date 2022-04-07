package main

import (
	"bufio"
	"os"
	"strings"
)

func rotateString(s string, goal string) bool {
	return len(s) == len(goal) && strings.Contains(s+s, goal)
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		arr := strings.Fields(input.Text())
		println(rotateString(arr[0], arr[1]))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了50.45%的用户
**/
