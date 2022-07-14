package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
	"unicode"
)

// https://leetcode.cn/problems/shortest-completing-word/
//------------------------Leetcode Problem 748------------------------
func shortestCompletingWord(licensePlate string, words []string) string {
	cnt := make([]int, 26)
	for _, x := range licensePlate {
		if unicode.IsLetter(x) {
			cnt[unicode.ToLower(x)-'a']++
		}
	}
	var ans string
next:
	for _, word := range words {
		c := make([]int, 26)
		for _, x := range word {
			c[x-'a']++
		}
		for i := 0; i < 26; i++ {
			if c[i] < cnt[i] {
				continue next
			}
		}
		if ans == "" || len(word) < len(ans) {
			ans = word
		}
	}
	return ans
}

//------------------------Leetcode Problem 748------------------------
/*
 * https://leetcode.cn/problems/shortest-completing-word/
 * 执行用时：4ms 在所有Go提交中击败了90.91%的用户
 * 占用内存：3.9MB 在所有Go提交中击败了87.88%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		licensePlate := scanner.Text()
		scanner.Scan()
		words := strings.Fields(scanner.Text())
		Printf("Output: %s\n", shortestCompletingWord(licensePlate, words))
	}
}
