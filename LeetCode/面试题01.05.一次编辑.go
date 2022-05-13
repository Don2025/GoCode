package main

import (
	"bufio"
	"os"
	"strings"
)

func oneEditAway(first string, second string) bool {
	m, n := len(first), len(second)
	if m < n {
		return oneEditAway(second, first)
	}
	if m-n > 1 {
		return false
	}
	for i := range second {
		if first[i] != second[i] {
			if m == n {
				return first[i+1:] == second[i+1:]
			}
			return first[i+1:] == second[i:]
		}
	}
	return true
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		arr := strings.Fields(input.Text())
		println(oneEditAway(arr[0], arr[1]))
	}
}

/*
 * 执行用时：4ms 在所有Go提交中击败了47.72%的用户
 * 占用内存：2.3MB 在所有Go提交中击败了63.86%的用户
**/
