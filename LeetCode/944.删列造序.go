package main

import (
	"bufio"
	"os"
	"strings"
)

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

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		strs := strings.Fields(input.Text())
		println(minDeletionSize(strs))
	}
}

/*
 * 执行用时：12ms 在所有Go提交中击败了34.09%的用户
 * 占用内存：6.3MB 在所有Go提交中击败了72.73%的用户
**/
