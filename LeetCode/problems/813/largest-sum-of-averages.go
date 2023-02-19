package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"math"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/largest-sum-of-averages/
//------------------------Leetcode Problem 813------------------------
func largestSumOfAverages(nums []int, k int) float64 {
	n := len(nums)
	prefix := make([]float64, n+1)
	for i, num := range nums {
		prefix[i+1] = prefix[i] + float64(num)
	}
	dp := make([]float64, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = prefix[i] / float64(i)
	}
	for i := 2; i <= k; i++ {
		for j := n; j >= i; j-- {
			for k := i - 1; k < j; k++ {
				dp[j] = math.Max(dp[j], dp[k]+(prefix[j]-prefix[k])/float64(j-k))
			}
		}
	}
	return dp[n]
}

//------------------------Leetcode Problem 813------------------------
/*
 * https://leetcode.cn/problems/largest-sum-of-averages/
 * 执行用时：4ms 在所有Go提交中击败了75.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了100.00%的用户
**/

func main() {
	scannner := bufio.NewScanner(os.Stdin)
	for scannner.Scan() {
		nums := utils.StringToInts(scannner.Text())
		scannner.Scan()
		k, _ := strconv.Atoi(scannner.Text())
		Printf("Output: %v\n", largestSumOfAverages(nums, k))
	}
}
