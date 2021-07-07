package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func countPairs(deliciousness []int) int {
	var ans int
	maxVal := deliciousness[0]
	for _, x := range deliciousness[1:] {
		maxVal = max(maxVal, x)
	}
	maxSum := maxVal * 2
	mp := make(map[int]int)
	for _, x := range deliciousness {
		for sum := 1; sum <= maxSum; sum <<= 1 {
			ans += mp[sum-x]
		}
		mp[x]++
	}
	return ans % 1000000007
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/* 暴力破解直接TLE
func countPairs(deliciousness []int) int {
	var ans int
	for i := 0; i < len(deliciousness); i++ {
		for j := i+1; j < len(deliciousness); j++ {
			sum := deliciousness[i]+deliciousness[j]
			if sum != 0 && (sum&(sum-1))==0 {
				ans++
				ans %= 1000000007
			}
		}
	}
	return ans
}
*/

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		println(countPairs(stringArrayToIntArray(strings.Fields(input.Text()))))
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
 * 执行用时：184ms 在所有Go提交中击败了53.45%的用户
 * 占用内存：8.9MB 在所有Go提交中击败了89.66%的用户
**/
