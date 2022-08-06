package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
)

// https://leetcode.cn/problems/string-matching-in-an-array/
//------------------------Leetcode Problem 1408------------------------
func stringMatching(words []string) []string {
	var ans []string
	for i := range words {
		for j := range words {
			if i != j && strings.Contains(words[j], words[i]) {
				ans = append(ans, words[i])
				break
			}
		}
	}
	return ans
}

//------------------------Leetcode Problem 1408------------------------
/*
 * https://leetcode.cn/problems/string-matching-in-an-array/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.2MB 在所有Go提交中击败了77.42%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		Printf("Output: %v\n", stringMatching(words))
	}
}
