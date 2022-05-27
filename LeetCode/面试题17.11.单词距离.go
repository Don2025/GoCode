package main

import (
	"bufio"
	"os"
	"strings"
)

func findClosest(words []string, word1 string, word2 string) int {
	ans := len(words)
	idx1, idx2 := -1, -1
	for i, word := range words {
		if word == word1 {
			idx1 = i
		} else if word == word2 {
			idx2 = i
		}
		if idx1 >= 0 && idx2 >= 0 {
			ans = min(ans, abs(idx2-idx1))
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(n int) int {
	if n < 0 {
		n = -n
	}
	return n
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		words := strings.Fields(input.Text())
		input.Scan()
		arr := strings.Fields(input.Text())
		println(findClosest(words, arr[0], arr[1]))
	}
}

/*
 * 执行用时：80ms 在所有Go提交中击败了37.14%的用户
 * 占用内存：11.3MB 在所有Go提交中击败了87.14%的用户
**/
