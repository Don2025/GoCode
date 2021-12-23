package main

import (
	"bufio"
	"os"
	"strings"
)

func checkInclusion(s1 string, s2 string) bool {
	n, m := len(s1), len(s2)
	if n > m {
		return false
	}
	var cnt1, cnt2 [26]int
	for i := range s1 {
		cnt1[s1[i]-'a']++
		cnt2[s2[i]-'a']++
	}
	if cnt1 == cnt2 {
		return true
	}
	for i := n; i < m; i++ {
		cnt2[s2[i]-'a']++
		cnt2[s2[i-n]-'a']--
		if cnt1 == cnt2 {
			return true
		}
	}
	return false
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		arr := strings.Fields(input.Text())
		s1, s2 := arr[0], arr[1]
		println(checkInclusion(s1, s2))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.6MB 在所有Go提交中击败了78.30%的用户
**/
