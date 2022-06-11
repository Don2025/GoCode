package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/0on3uN/
// ------------------------剑指 Offer II Problem 87------------------------
func restoreIpAddresses(s string) (ans []string) {
	const SEG_COUNT = 4
	segments := make([]int, SEG_COUNT)
	var dfs func(int, int)
	dfs = func(segId int, segStart int) {
		if segId == SEG_COUNT {
			if segStart == len(s) {
				ipAddr := ""
				for i := 0; i < SEG_COUNT; i++ {
					ipAddr += strconv.Itoa(segments[i])
					if i != SEG_COUNT-1 {
						ipAddr += "."
					}
				}
				ans = append(ans, ipAddr)
			}
			return
		}
		if segStart == len(s) {
			return
		}
		if s[segStart] == '0' {
			segments[segId] = 0
			dfs(segId+1, segStart+1)
		}
		addr := 0
		for segEnd := segStart; segEnd < len(s); segEnd++ {
			addr = addr*10 + int(s[segEnd]-'0')
			if addr > 0 && addr <= 0xFF {
				segments[segId] = addr
				dfs(segId+1, segEnd+1)
			} else {
				break
			}
		}
	}
	dfs(0, 0)
	return
}

// ------------------------剑指 Offer II Problem 87------------------------
/*
 * https://leetcode.cn/problems/0on3uN/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了97.16%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		Printf("%v\n", restoreIpAddresses(scanner.Text()))
	}
}
