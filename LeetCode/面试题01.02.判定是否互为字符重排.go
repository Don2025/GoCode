package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
)

func CheckPermutation(s1 string, s2 string) bool {
	var c1, c2 [128]int
	for _, ch := range s1 {
		c1[ch]++
	}
	for _, ch := range s2 {
		c2[ch]++
	}
	return c1 == c2
}

/*
 * https://leetcode.cn/problems/check-permutation-lcci/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了38.64%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		strs := strings.Fields(scanner.Text())
		Printf("Output: %t\n", CheckPermutation(strs[0], strs[1]))
	}
}
