package main

import (
	"bufio"
	. "fmt"
	"os"
)

// https://leetcode.cn/problems/partition-labels/
//------------------------Leetcode Problem 763------------------------
func partitionLabels(s string) []int {
	var ans []int
	m := [26]int{}
	for i, x := range s {
		m[x-'a'] = i
	}
	var start, end int
	for i, x := range s {
		end = max(end, m[x-'a'])
		if i == end {
			ans = append(ans, end-start+1)
			start = i + 1
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//------------------------Leetcode Problem 763------------------------
/*
 * https://leetcode.cn/problems/partition-labels/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2MB 在所有Go提交中击败了100.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		Printf("Output: %v\n", partitionLabels(scanner.Text()))
	}
}
