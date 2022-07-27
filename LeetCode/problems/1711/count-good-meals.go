package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

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

// https://leetcode.cn/problems/count-good-meals/
//------------------------Leetcode Problem 1711------------------------
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

//------------------------Leetcode Problem 1711------------------------
/*
 * https://leetcode.cn/problems/count-good-meals/
 * 执行用时：184ms 在所有Go提交中击败了53.45%的用户
 * 占用内存：8.9MB 在所有Go提交中击败了89.66%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		deliciousness := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", countPairs(deliciousness))
	}
}
