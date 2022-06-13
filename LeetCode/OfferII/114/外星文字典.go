package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
)

// https://leetcode.cn/problems/Jf1JuT/
// ------------------------剑指 Offer II Problem 114------------------------
func alienOrder(words []string) string {
	g := make(map[byte][]byte)
	for _, x := range words[0] {
		g[byte(x)] = nil
	}
next:
	for i := 1; i < len(words); i++ {
		s, t := words[i-1], words[i]
		for _, x := range t {
			g[byte(x)] = g[byte(x)]
		}
		for j := 0; j < len(s) && j < len(t); j++ {
			if s[j] != t[j] {
				g[s[j]] = append(g[s[j]], t[j])
				continue next
			}
		}
		if len(s) > len(t) {
			return ""
		}
	}
	const visiting, visited = 1, 2
	order := make([]byte, len(g))
	i := len(g) - 1
	state := make(map[byte]int)
	var dfs func(u byte) bool
	dfs = func(u byte) bool {
		state[u] = visiting
		for _, v := range g[u] {
			if state[v] == 0 {
				if !dfs(v) {
					return false
				}
			} else if state[v] == visiting {
				return false
			}
		}
		order[i] = u
		i--
		state[u] = visited
		return true
	}
	for u := range g {
		if state[u] == 0 && !dfs(u) {
			return ""
		}
	}
	return string(order)
}

// ------------------------剑指 Offer II Problem 114------------------------
/*
 * https://leetcode.cn/problems/Jf1JuT/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.2MB 在所有Go提交中击败了34.61%的用户
**/
func main() {
	input := bufio.NewScanner(os.Stdin)
	Printf("Input a line of strings separated by space:")
	for input.Scan() {
		words := strings.Fields(input.Text())
		Printf("Output: %v\n", alienOrder(words))
		Printf("Input a line of strings separated by space:")
	}
}
