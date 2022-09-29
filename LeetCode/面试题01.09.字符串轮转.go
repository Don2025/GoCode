package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
)

func isFlipedString(s1 string, s2 string) bool {
	return len(s1) == len(s2) && strings.Contains(s1+s1, s2)
}

/*
 * https://leetcode.cn/problems/string-rotation-lcci/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.9MB 在所有Go提交中击败了50.36%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		strs := strings.Fields(scanner.Text())
		Printf("Output: %t\n", isFlipedString(strs[0], strs[1]))
	}
}
