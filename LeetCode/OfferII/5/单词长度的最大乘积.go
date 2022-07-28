package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
)

// https://leetcode.cn/problems/aseY1I/
// ------------------------剑指 Offer II Problem 5------------------------
func maxProduct(words []string) int {
	masks := make([]int, len(words))
	for i, word := range words {
		for _, x := range word {
			masks[i] |= 1 << (x - 'a')
		}
	}
	var ans int
	for i, x := range masks {
		for j, y := range masks[:i] {
			if t := len(words[i]) * len(words[j]); x&y == 0 && t > ans {
				ans = t
			}
		}
	}
	return ans
}

// ------------------------剑指 Offer II Problem 5------------------------
/*
 * https://leetcode.cn/problems/aseY1I/
 * 执行用时：8ms 在所有Go提交中击败了97.96%的用户
 * 占用内存：6.3MB 在所有Go提交中击败了78.48%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		Printf("Output: %v\n", maxProduct(words))
	}
}
