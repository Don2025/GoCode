package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/matchsticks-to-square/
//------------------------Leetcode Problem 473------------------------
func makesquare(matchsticks []int) bool {
	var totalLen int
	for _, l := range matchsticks {
		totalLen += l
	}
	if totalLen%4 != 0 {
		return false
	}
	tLen := totalLen / 4
	dp := make([]int, 1<<len(matchsticks))
	for i := 1; i < len(dp); i++ {
		dp[i] = -1
	}
	for i := 1; i < len(dp); i++ {
		for k, v := range matchsticks {
			if i>>k&1 == 0 {
				continue
			}
			j := i &^ (1 << k)
			if dp[j] >= 0 && dp[j]+v <= tLen {
				dp[i] = (dp[j] + v) % tLen
				break
			}
		}
	}
	return dp[len(dp)-1] == 0
}

//------------------------Leetcode Problem 473------------------------
/*
 * https://leetcode.cn/problems/matchsticks-to-square/
 * 执行用时：76ms 在所有Go提交中击败了32.03%的用户
 * 占用内存：6.7MB 在所有Go提交中击败了5.47%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		matchsticks := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", makesquare(matchsticks))
	}
}
