package main

import (
	"bufio"
	"os"
	"strings"
)

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

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		words := strings.Fields(input.Text())
		println(maxProduct(words))
	}
}

/*
 * 执行用时：8ms 在所有Go提交中击败了97.96%的用户
 * 占用内存：6.3MB 在所有Go提交中击败了78.48%的用户
**/
