package main

import (
	"bufio"
	"os"
	"strings"
	"unicode"
)

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

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		licensePlate := input.Text()
		input.Scan()
		words := strings.Fields(input.Text())
		println(shortestCompletingWord(licensePlate, words))
	}
}

/*
 * 执行用时：4ms 在所有Go提交中击败了90.91%的用户
 * 占用内存：3.9MB 在所有Go提交中击败了87.88%的用户
**/
