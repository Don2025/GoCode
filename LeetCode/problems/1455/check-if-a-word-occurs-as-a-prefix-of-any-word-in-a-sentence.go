package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
)

// https://leetcode.cn/problems/check-if-a-word-occurs-as-a-prefix-of-any-word-in-a-sentence/
//------------------------Leetcode Problem 1455------------------------
func isPrefixOfWord(sentence string, searchWord string) int {
	for i, index := 0, 1; i < len(sentence); i++ {
		start := i
		for i < len(sentence) && sentence[i] != ' ' {
			i++
		}
		end := i
		if strings.HasPrefix(sentence[start:end], searchWord) {
			return index
		}
		index++
	}
	return -1
}

//------------------------Leetcode Problem 1455------------------------
/*
 * https://leetcode.cn/problems/check-if-a-word-occurs-as-a-prefix-of-any-word-in-a-sentence/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了12.50%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		sentence := scanner.Text()
		scanner.Scan()
		searchWord := scanner.Text()
		Printf("Output: %v\n", isPrefixOfWord(sentence, searchWord))
	}
}
