package main

import (
	"bufio"
	"os"
	"path/filepath"
)

func simplifyPath(path string) string {
	return filepath.Clean(path)
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		println(simplifyPath(input.Text()))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.6MB 在所有Go提交中击败了96.14%的用户
**/
