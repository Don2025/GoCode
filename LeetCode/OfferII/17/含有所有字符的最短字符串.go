package main

import (
	"bufio"
	"math"
	"os"
	"strings"
)

// https://leetcode.cn/problems/M1oyTv/
// ------------------------剑指 Offer II Problem 17------------------------
func minWindow(s string, t string) string {
	sCnt, tCnt := map[byte]int{}, map[byte]int{}
	for i := range t {
		tCnt[t[i]]++
	}
	sLen, n := len(s), math.MaxInt32
	ansL, ansR := -1, -1
	check := func() bool {
		for k, v := range tCnt {
			if sCnt[k] < v {
				return false
			}
		}
		return true
	}
	for l, r := 0, 0; r < sLen; r++ {
		if r < sLen && tCnt[s[r]] > 0 {
			sCnt[s[r]]++
		}
		for check() && l <= r {
			if r-l+1 < n {
				n = r - l + 1
				ansL, ansR = l, l+n
			}
			if _, ok := tCnt[s[l]]; ok {
				sCnt[s[l]]--
			}
			l++
		}
	}
	if ansL == -1 {
		return ""
	}
	return s[ansL:ansR]
}

// ------------------------剑指 Offer II Problem 17------------------------
/*
 * https://leetcode.cn/problems/M1oyTv/
 * 执行用时：108ms 在所有Go提交中击败了22.57%的用户
 * 占用内存：2.7MB 在所有Go提交中击败了52.33%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		arr := strings.Fields(scanner.Text())
		println(minWindow(arr[0], arr[1]))
	}
}
