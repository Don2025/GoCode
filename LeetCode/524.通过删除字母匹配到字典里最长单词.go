package main

import (
	"bufio"
	"os"
	"sort"
	"strings"
)

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

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		s := input.Text()
		input.Scan()
		dictionary := strings.Fields(input.Text())
		println(findLongestWord(s, dictionary))
	}
}

/*
 * 执行用时：20ms 在所有Go提交中击败了28.28%的用户
 * 占用内存：7.8MB 在所有Go提交中击败了34.48%的用户
**/
