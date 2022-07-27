package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
)

// https://leetcode.cn/problems/maximum-length-of-a-concatenated-string-with-unique-characters/
//------------------------Leetcode Problem 1239------------------------
func maxLength(arr []string) int {
	var backtrace func(int, int) int
	backtrace = func(idx, bit int) int {
		var ans int
		if idx == len(arr) {
			return ans
		}
		tmp := bit
		for _, x := range arr[idx] {
			t := 1 << (x - 'a')
			if bit&t > 0 {
				return backtrace(idx+1, tmp)
			}
			bit |= t
		}
		return max(len(arr[idx])+backtrace(idx+1, bit), backtrace(idx+1, tmp))
	}
	return backtrace(0, 0)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		Printf("Output: %v\n", maxLength(strings.Split(scanner.Text(), " ")))
	}
}

/*
 * https://leetcode.cn/problems/maximum-length-of-a-concatenated-string-with-unique-characters/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2MB 在所有Go提交中击败了86.11%的用户
**/
