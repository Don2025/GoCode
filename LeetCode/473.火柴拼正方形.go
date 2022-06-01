package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

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

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		matchsticks := stringArrayToIntArray(strings.Fields(input.Text()))
		println(makesquare(matchsticks))
	}
}

func stringArrayToIntArray(strings []string) []int {
	var arr []int
	for _, x := range strings {
		n, _ := strconv.Atoi(x)
		arr = append(arr, n)
	}
	return arr
}

/*
 * 执行用时：76ms 在所有Go提交中击败了32.03%的用户
 * 占用内存：6.7MB 在所有Go提交中击败了5.47%的用户
**/
