package main

import (
	"bufio"
	. "fmt"
	"os"
)

// https://leetcode.cn/problems/replace-all-s-to-avoid-consecutive-repeating-characters/
//------------------------Leetcode Problem 1576------------------------
func modifyString(s string) string {
	arr, n := []int32(s), len(s)
	for i := range arr {
		if arr[i] == '?' {
			for ch := 'a'; ch <= 'c'; ch++ {
				if (i > 0 && arr[i-1] == ch) || (i < n-1 && arr[i+1] == ch) {
					continue
				}
				arr[i] = ch
				break
			}
		}
	}
	return string(arr)
}

//------------------------Leetcode Problem 1576------------------------
/*
 * https://leetcode.cn/problems/replace-all-s-to-avoid-consecutive-repeating-characters/
 * 执行用时：4ms 在所有Go提交中击败了44.83%的用户
 * 占用内存：2.2MB 在所有Go提交中击败了10.34%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		Printf("Output: %v\n", modifyString(scanner.Text()))
	}
}
