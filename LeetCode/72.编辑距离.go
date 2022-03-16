package main

import (
	"bufio"
	"os"
)

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

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		s1 := input.Text()
		input.Scan()
		s2 := input.Text()
		println(minDistance(s1, s2))
	}
}

/*
 * 执行用时：4ms 在所有Go提交中击败了79.04%的用户
 * 占用内存：5.4MB 在所有Go提交中击败了91.49%的用户
**/
