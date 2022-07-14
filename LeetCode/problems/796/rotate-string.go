package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
)

// https://leetcode.cn/problems/rotate-string/
//------------------------Leetcode Problem 796------------------------
func rotateString(s string, goal string) bool {
	return len(s) == len(goal) && strings.Contains(s+s, goal)
}

//------------------------Leetcode Problem 796------------------------
/*
 * https://leetcode.cn/problems/rotate-string/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了50.45%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		arr := strings.Fields(scanner.Text())
		Printf("Output: %v\n", rotateString(arr[0], arr[1]))
	}
}
