package main

import (
	"bufio"
	"os"
	"strings"
)

func interpret(command string) string {
	command = strings.ReplaceAll(command, "()", "o")
	command = strings.ReplaceAll(command, "(al)", "al")
	return command
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		command := input.Text()
		println(interpret(command))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.8MB 在所有Go提交中击败了86.76%的用户
**/
