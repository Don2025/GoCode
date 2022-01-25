package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func findSubstring(s string, words []string) []int {
	var ans []int
	nw := len(words)
	if nw == 0 {
		return ans
	}
	ns, k := len(s), len(words[0])
	if ns < nw*k {
		return ans
	}
	m := map[string]int{}
	for _, x := range words {
		m[x]++
	}
	for i := 0; i < k; i++ {
		var cnt int
		tm := map[string]int{}
		for l, r := i, i; r <= ns-k; r += k {
			word := s[r : r+k]
			if x, ok := m[word]; ok {
				for tm[word] >= x {
					tm[s[l:l+k]]--
					cnt--
					l += k
				}
				tm[word]++
				cnt++
			} else {
				for l < r {
					tm[s[l:l+k]]--
					cnt--
					l += k
				}
				l += k
			}
			if cnt == nw {
				ans = append(ans, l)
			}
		}
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		s := input.Text()
		input.Scan()
		words := strings.Fields(input.Text())
		ans := findSubstring(s, words)
		for _, x := range ans {
			fmt.Printf("%d ", x)
		}
		fmt.Println()
	}
}

/*
 * 执行用时：4ms 在所有Go提交中击败了98.85%的用户
 * 占用内存：3.2MB 在所有Go提交中击败了94.07%的用户
**/
