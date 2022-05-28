package main

import (
	"bufio"
	"os"
	"strings"
)

func isAlienSorted(words []string, order string) bool {
	var index [26]int
	for i, x := range order {
		index[x-'a'] = i
	}
	for i := 1; i < len(words); i++ {
		var valid bool
		for j := 0; j < len(words[i-1]) && j < len(words[i]); j++ {
			prev := index[words[i-1][j]-'a']
			cur := index[words[i][j]-'a']
			if prev < cur {
				valid = true
				break
			} else if prev > cur {
				return false
			}
		}
		if !valid {
			return len(words[i-1]) <= len(words[i])
		}
	}
	return true
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		words := strings.Fields(input.Text())
		input.Scan()
		println(isAlienSorted(words, input.Text()))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.4MB 在所有Go提交中击败了79.91%的用户
**/
