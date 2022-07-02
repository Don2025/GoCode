package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
	"strings"
)

// https://leetcode.cn/problems/longest-word-in-dictionary-through-deleting/
//------------------------Leetcode Problem 524------------------------
func findLongestWord(s string, dictionary []string) string {
	sort.Slice(dictionary, func(i, j int) bool {
		return len(dictionary[i]) > len(dictionary[j]) || len(dictionary[i]) == len(dictionary[j]) && dictionary[i] < dictionary[j]
	})
	for _, x := range dictionary {
		var cnt int
		for i := range s {
			if s[i] == x[cnt] {
				cnt++
				if cnt == len(x) {
					return x
				}
			}
		}
	}
	return ""
}

//------------------------Leetcode Problem 524------------------------
/*
 * https://leetcode.cn/problems/longest-word-in-dictionary-through-deleting/
 * 执行用时：20ms 在所有Go提交中击败了28.28%的用户
 * 占用内存：7.8MB 在所有Go提交中击败了34.48%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()
		scanner.Scan()
		dictionary := strings.Fields(scanner.Text())
		Printf("Output: %v\n", findLongestWord(s, dictionary))
	}
}
