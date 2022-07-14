package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
)

// https://leetcode.cn/problems/unique-morse-code-words/
//------------------------Leetcode Problem 804------------------------
func uniqueMorseRepresentations(words []string) int {
	morse := []string{
		".-", "-...", "-.-.", "-..", ".", "..-.", "--.",
		"....", "..", ".---", "-.-", ".-..", "--", "-.",
		"---", ".--.", "--.-", ".-.", "...", "-", "..-",
		"...-", ".--", "-..-", "-.--", "--..",
	}
	set := map[string]bool{}
	for _, word := range words {
		trans := &strings.Builder{}
		for _, ch := range word {
			trans.WriteString(morse[ch-'a'])
		}
		set[trans.String()] = true
	}
	return len(set)
}

//------------------------Leetcode Problem 804------------------------
/*
 * https://leetcode.cn/problems/unique-morse-code-words/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.2MB 在所有Go提交中击败了99.93%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		Printf("Output: %v\n", uniqueMorseRepresentations(words))
	}
}
