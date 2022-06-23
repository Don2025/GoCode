package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
)

// https://leetcode.cn/problems/substring-with-concatenation-of-all-words/
//------------------------Leetcode Problem 30------------------------
func findSubstring(s string, words []string) (ans []int) {
	ls, m, n := len(s), len(words), len(words[0])
	for i := 0; i < n && i+m*n <= ls; i++ {
		differ := map[string]int{}
		for j := 0; j < m; j++ {
			differ[s[i+j*n:i+(j+1)*n]]++
		}
		for _, word := range words {
			differ[word]--
			if differ[word] == 0 {
				delete(differ, word)
			}
		}
		for start := i; start < ls-m*n+1; start += n {
			if start != i {
				word := s[start+(m-1)*n : start+m*n]
				differ[word]++
				if differ[word] == 0 {
					delete(differ, word)
				}
				word = s[start-n : start]
				differ[word]--
				if differ[word] == 0 {
					delete(differ, word)
				}
			}
			if len(differ) == 0 {
				ans = append(ans, start)
			}
		}
	}
	return
}

//------------------------Leetcode Problem 30------------------------
/*
 * https://leetcode.cn/problems/substring-with-concatenation-of-all-words/
 * 执行用时：12ms 在所有Go提交中击败了72.86%的用户
 * 占用内存：6.6MB 在所有Go提交中击败了42.32%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	Printf("Input a line of string:")
	for scanner.Scan() {
		s := scanner.Text()
		Printf("Input a line of strings separated by space:")
		scanner.Scan()
		words := strings.Fields(scanner.Text())
		Printf("Output: %v\n", findSubstring(s, words))
		Printf("Input a line of string:")
	}
}
