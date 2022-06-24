package main

import (
	"bufio"
	. "fmt"
	"os"
)

//https://leetcode.cn/problems/edit-distance/
//------------------------Leetcode Problem 72------------------------
func minDistance(word1 string, word2 string) int {
	n1, n2 := len(word1), len(word2)
	cost := make([][]int, n1+1)
	for i := 0; i <= n1; i++ {
		cost[i] = make([]int, n2+1)
		cost[i][0] = i
	}
	for i := 0; i <= n2; i++ {
		cost[0][i] = i
	}
	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if word1[i-1] == word2[j-1] {
				cost[i][j] = cost[i-1][j-1]
			} else {
				cost[i][j] = 1 + min(cost[i-1][j-1], min(cost[i][j-1], cost[i-1][j]))
			}
		}
	}
	return cost[n1][n2]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//------------------------Leetcode Problem 72------------------------
/*
 * https://leetcode.cn/problems/edit-distance/
 * 执行用时：4ms 在所有Go提交中击败了79.04%的用户
 * 占用内存：5.4MB 在所有Go提交中击败了91.49%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s1 := scanner.Text()
		scanner.Scan()
		s2 := scanner.Text()
		Printf("Output: %v\n", minDistance(s1, s2))
	}
}
