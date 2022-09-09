package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
)

// https://leetcode.cn/problems/crawler-log-folder/
//------------------------Leetcode Problem 1598------------------------
func minOperations(logs []string) int {
	var cnt int
	for _, log := range logs {
		if strings.EqualFold(log, "./") {
			continue
		} else if strings.EqualFold(log, "../") {
			if cnt > 0 {
				cnt--
			}
		} else {
			cnt++
		}
	}
	return cnt
}

//------------------------Leetcode Problem 1598------------------------
/*
 * https://leetcode.cn/problems/crawler-log-folder/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.9MB 在所有Go提交中击败了100.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		logs := strings.Fields(scanner.Text())
		Printf("Output: %v\n", minOperations(logs))
	}
}
