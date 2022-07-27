package main

import (
	"bufio"
	"os"
)

// https://leetcode.cn/problems/check-if-one-string-swap-can-make-strings-equal/
//------------------------Leetcode Problem 1790------------------------
func areAlmostEqual(s1 string, s2 string) bool {
	var arr []int
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			arr = append(arr, i)
		}
	}
	if len(arr) == 0 {
		return true
	} else if len(arr) == 2 {
		if s1[arr[0]] == s2[arr[1]] && s1[arr[1]] == s2[arr[0]] {
			return true
		}
	}
	return false
}

//------------------------Leetcode Problem 1790------------------------
/*
 * https://leetcode.cn/problems/check-if-one-string-swap-can-make-strings-equal/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了41.67%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s1 := scanner.Text()
		scanner.Scan()
		s2 := scanner.Text()
		println(areAlmostEqual(s1, s2))
	}
}
