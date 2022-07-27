package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
)

// https://leetcode.cn/problems/goal-parser-interpretation/
//------------------------Leetcode Problem 1678------------------------
func interpret(command string) string {
	command = strings.ReplaceAll(command, "()", "o")
	command = strings.ReplaceAll(command, "(al)", "al")
	return command
}

//------------------------Leetcode Problem 1678------------------------
/*
 * https://leetcode.cn/problems/goal-parser-interpretation/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.8MB 在所有Go提交中击败了86.76%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		command := scanner.Text()
		Printf("Output: %v\n", interpret(command))
	}
}
