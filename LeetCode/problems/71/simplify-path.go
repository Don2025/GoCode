package main

import (
	"bufio"
	. "fmt"
	"os"
	"path/filepath"
)

// https://leetcode.cn/problems/simplify-path/
//------------------------Leetcode Problem 71------------------------
func simplifyPath(path string) string {
	return filepath.Clean(path)
}

//------------------------Leetcode Problem 71------------------------
/*
 * https://leetcode.cn/problems/simplify-path/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.6MB 在所有Go提交中击败了96.14%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		Printf("Output: %v\n", simplifyPath(scanner.Text()))
	}
}
