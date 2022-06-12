package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
)

// https://leetcode.cn/problems/zlDJc7
// ------------------------剑指 Offer II Problem 109------------------------
func openLock(deadends []string, target string) int {
	const start = "0000"
	if target == start {
		return 0
	}
	dead := make(map[string]bool)
	for _, x := range deadends {
		dead[x] = true
	}
	if dead[start] {
		return -1
	}
	get := func(status string) (ret []string) {
		s := []byte(status)
		for i, b := range s {
			s[i] = b - 1
			if s[i] < '0' {
				s[i] = '9'
			}
			ret = append(ret, string(s))
			s[i] = b + 1
			if s[i] > '9' {
				s[i] = '0'
			}
			ret = append(ret, string(s))
			s[i] = b
		}
		return
	}
	type pair struct {
		status string
		step   int
	}
	q := []pair{{start, 0}}
	seen := map[string]bool{start: true}
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		for _, nxt := range get(p.status) {
			if !seen[nxt] && !dead[nxt] {
				if nxt == target {
					return p.step + 1
				}
				seen[nxt] = true
				q = append(q, pair{nxt, p.step + 1})
			}
		}
	}
	return -1
}

// ------------------------剑指 Offer II Problem 109------------------------
/*
 * https://leetcode.cn/problems/2bCMpM/
 * 执行用时：88ms 在所有Go提交中击败了68.40%的用户
 * 占用内存：7.4MB 在所有Go提交中击败了67.44%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	Printf("Input a line of digital strings separated by space:")
	for scanner.Scan() {
		strs := strings.Fields(scanner.Text())
		Printf("Input a line of digital string(target):")
		scanner.Scan()
		target := scanner.Text()
		Printf("Output: %v\n", openLock(strs, target))
		Printf("Input a line of digital strings separated by space:")
	}
}
