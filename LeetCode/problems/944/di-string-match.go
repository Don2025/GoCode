package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
)

// https://leetcode.cn/problems/delete-columns-to-make-sorted/
//------------------------Leetcode Problem 944------------------------
func minDeletionSize(strs []string) int {
	var cnt int
	for i := 0; i < len(strs[0]); i++ {
		for j := 0; j < len(strs)-1; j++ {
			if strs[j][i] > strs[j+1][i] {
				cnt++
				break
			}
		}
	}
	return cnt
}

//------------------------Leetcode Problem 944------------------------
/*
 * https://leetcode.cn/problems/di-string-match/
 * 执行用时：12ms 在所有Go提交中击败了34.09%的用户
 * 占用内存：6.3MB 在所有Go提交中击败了72.73%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		strs := strings.Fields(scanner.Text())
		Printf("Output: %d\n", minDeletionSize(strs))
	}
}
