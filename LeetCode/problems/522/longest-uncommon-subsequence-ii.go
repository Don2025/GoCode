package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
	"strings"
)

// https://leetcode.cn/problems/longest-uncommon-subsequence-ii/
//------------------------Leetcode Problem 522------------------------
func findLUSlength(strs []string) int {
	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) > len(strs[j])
	})
	n := len(strs)
	for i := 0; i < n; i++ {
		flag := true
		for j := 0; j < n && len(strs[i]) <= len(strs[j]); j++ {
			if i == j {
				continue
			}
			var p1, p2 int
			for p1 < len(strs[i]) && p2 < len(strs[j]) {
				if strs[i][p1] != strs[j][p2] {
					p2++
				} else {
					p1++
					p2++
				}
			}
			if p1 == len(strs[i]) {
				flag = false
				break
			}
		}
		if flag {
			return len(strs[i])
		}
	}
	return -1
}

//------------------------Leetcode Problem 522------------------------
/*
 * https://leetcode.cn/problems/longest-uncommon-subsequence-ii/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了70.77%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		strs := strings.Fields(scanner.Text())
		Printf("Output: %v\n", findLUSlength(strs))
	}
}
