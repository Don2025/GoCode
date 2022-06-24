package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
)

//https://leetcode.cn/problems/length-of-last-word/
//------------------------Leetcode Problem 58------------------------
func lengthOfLastWord(s string) int {
	arr := strings.Fields(s)
	return len(arr[len(arr)-1])
}

//------------------------Leetcode Problem 58------------------------
/*
 * https://leetcode.cn/problems/length-of-last-word/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.1MB 在所有Go提交中击败了28.43%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		Printf("%v\n", lengthOfLastWord(scanner.Text()))
	}
}
