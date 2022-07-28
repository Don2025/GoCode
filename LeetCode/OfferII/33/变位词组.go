package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
	"strings"
)

// https://leetcode.cn/problems/sfvd7V/
// ------------------------剑指 Offer II Problem 33------------------------
func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
		sortedStr := string(s)
		m[sortedStr] = append(m[sortedStr], str)
	}
	ans := make([][]string, 0, len(m))
	for _, row := range m {
		ans = append(ans, row)
	}
	return ans
}

// ------------------------剑指 Offer II Problem 33------------------------
/*
 * https://leetcode.cn/problems/sfvd7V/
 * 执行用时：20ms 在所有Go提交中击败了74.31%的用户
 * 占用内存：8.2MB 在所有Go提交中击败了70.49%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		strs := strings.Fields(scanner.Text())
		Printf("Output: %v\n", groupAnagrams(strs))
	}
}
