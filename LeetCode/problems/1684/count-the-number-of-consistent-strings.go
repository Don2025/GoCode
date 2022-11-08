package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
)

// https://leetcode.cn/problems/count-the-number-of-consistent-strings/
//------------------------Leetcode Problem 1684------------------------
func countConsistentStrings(allowed string, words []string) (cnt int) {
	f := func(s string) (mask int) {
		for _, c := range s {
			mask |= 1 << (c - 'a')
		}
		return
	}
	mask := f(allowed)
	for _, word := range words {
		if mask|f(word) == mask {
			cnt++
		}
	}
	return
}

//------------------------Leetcode Problem 1684------------------------
/*
 * https://leetcode.cn/problems/count-the-number-of-consistent-strings/
 * 执行用时：20ms 在所有Go提交中击败了97.92%的用户
 * 占用内存：6.6MB 在所有Go提交中击败了95.83%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		allowed := scanner.Text()
		scanner.Scan()
		words := strings.Fields(scanner.Text())
		Printf("Output: %v\n", countConsistentStrings(allowed, words))
	}
}
