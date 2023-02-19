package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
)

// https://leetcode.cn/problems/custom-sort-string/
//------------------------Leetcode Problem 791------------------------
func customSortString(order string, s string) string {
	cnt := make([]int, 26)
	for _, c := range s {
		cnt[c-'a']++
	}
	var sb strings.Builder
	for _, ch := range order {
		for cnt[ch-'a'] > 0 {
			sb.WriteByte(byte(ch))
			cnt[ch-'a']--
		}
	}
	for i, n := range cnt {
		if n > 0 {
			sb.WriteString(strings.Repeat(string('a'+i), n))
		}
	}
	return sb.String()
}

//------------------------Leetcode Problem 791------------------------
/*
 * https://leetcode.cn/problems/custom-sort-string/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了44.74%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		arr := strings.Fields(scanner.Text())
		Printf("Output: %v\n", customSortString(arr[0], arr[1]))
	}
}
