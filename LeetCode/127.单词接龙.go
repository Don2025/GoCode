package main

import (
	"bufio"
	"os"
	"strings"
)

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordMap := map[string]bool{}
	for _, w := range wordList {
		wordMap[w] = true
	}
	queue := []string{beginWord}
	level := 1
	for len(queue) != 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			word := queue[0]
			queue = queue[1:]
			if word == endWord {
				return level
			}
			for j := 0; j < len(word); j++ {
				for k := 'a'; k <= 'z'; k++ {
					newWord := word[:j] + string(k) + word[j+1:]
					if wordMap[newWord] {
						queue = append(queue, newWord)
						wordMap[newWord] = false
					}
				}
			}
		}
		level++
	}
	return 0
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		arr := strings.Fields(input.Text())
		beginWord, endWord := arr[0], arr[1]
		input.Scan()
		arr = strings.Fields(input.Text())
		println(ladderLength(beginWord, endWord, arr))
	}
}

/*
 * 执行用时：168ms 在所有Go提交中击败了31.63%的用户
 * 占用内存：7.1MB 在所有Go提交中击败了54.41%的用户
**/
