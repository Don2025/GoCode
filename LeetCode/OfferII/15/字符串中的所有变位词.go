package main

import (
	"bufio"
	. "fmt"
	"os"
)

// https://leetcode.cn/problems/VabMRr/
// ------------------------剑指 Offer II Problem 15------------------------
func findAnagrams(s string, p string) []int {
	var ans []int
	ns, np := len(s), len(p)
	if ns < np {
		return ans
	}
	var cnts, cntp [26]int
	for i := range p {
		cnts[s[i]-'a']++
		cntp[p[i]-'a']++
	}
	if cnts == cntp {
		ans = append(ans, 0)
	}
	for i := range s[:ns-np] {
		cnts[s[i]-'a']--
		cnts[s[i+np]-'a']++
		if cnts == cntp {
			ans = append(ans, i+1)
		}
	}
	return ans
}

// ------------------------剑指 Offer II Problem 15------------------------
/*
 * https://leetcode.cn/problems/VabMRr/
 * 执行用时：4ms 在所有Go提交中击败了90.18%的用户
 * 占用内存：4.9MB 在所有Go提交中击败了62.58%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var s, p string
		Sscanf(scanner.Text(), "%s %s", &s, &p)
		Printf("Output: %v\n", findAnagrams(s, p))
	}
}
