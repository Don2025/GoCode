package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
)

// https://leetcode.cn/problems/swap-adjacent-in-lr-string/
//------------------------Leetcode Problem 777------------------------
func canTransform(start, end string) bool {
	i, j, n := 0, 0, len(start)
	for i < n && j < n {
		for i < n && start[i] == 'X' {
			i++
		}
		for j < n && end[j] == 'X' {
			j++
		}
		if i < n && j < n {
			if start[i] != end[j] {
				return false
			}
			c := start[i]
			if c == 'L' && i < j || c == 'R' && i > j {
				return false
			}
			i++
			j++
		}
	}
	for i < n {
		if start[i] != 'X' {
			return false
		}
		i++
	}
	for j < n {
		if end[j] != 'X' {
			return false
		}
		j++
	}
	return true
}

//------------------------Leetcode Problem 777------------------------
/*
 * https://leetcode.cn/problems/swap-adjacent-in-lr-string/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.7MB 在所有Go提交中击败了100.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		strs := strings.Fields(scanner.Text())
		start, end := strs[0], strs[1]
		Printf("Output: %t\n", canTransform(start, end))
	}
}
