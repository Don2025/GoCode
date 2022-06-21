package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
)

// https://leetcode.cn/problems/defanging-an-ip-address/
//------------------------Leetcode Problem 1108------------------------
func defangIPaddr(address string) string {
	return strings.ReplaceAll(address, ".", "[.]")
}

//------------------------Leetcode Problem 1108------------------------
/*
 * https://leetcode.cn/problems/defanging-an-ip-address/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.8MB 在所有Go提交中击败了73.57%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	Printf("Input a line of string:")
	for scanner.Scan() {
		Printf("Output: %v\n", defangIPaddr(scanner.Text()))
		Printf("Input a line of string:")
	}
}
