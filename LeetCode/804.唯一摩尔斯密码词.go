package main

import (
	"bufio"
	"os"
	"strings"
)

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

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		words := strings.Fields(input.Text())
		println(uniqueMorseRepresentations(words))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.2MB 在所有Go提交中击败了99.93%的用户
**/
