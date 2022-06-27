package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
)

// https://leetcode.cn/problems/maximum-product-of-word-lengths/
//------------------------Leetcode Problem 318------------------------
func maxProduct(words []string) int {
	var ans int
	for i := 0; i < len(words)-1; i++ {
		for j := i + 1; j < len(words); j++ {
			if !strings.ContainsAny(words[i], words[j]) {
				ans = max(ans, len(words[i])*len(words[j]))
			}
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

//------------------------Leetcode Problem 318------------------------
/*
 * https://leetcode.cn/problems/maximum-product-of-word-lengths/
 * 执行用时：176ms 在所有Go提交中击败了22.38%的用户
 * 占用内存：6.3MB 在所有Go提交中击败了99.30%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		Printf("Output: %v\n", maxProduct(strings.Fields(scanner.Text())))
	}
}
