package main

import (
	"bufio"
	"os"
	"strings"
)

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

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		println(maxProduct(strings.Fields(input.Text())))
	}
}

/*
 * 执行用时：176ms 在所有Go提交中击败了22.38%的用户
 * 占用内存：6.3MB 在所有Go提交中击败了99.30%的用户
**/
