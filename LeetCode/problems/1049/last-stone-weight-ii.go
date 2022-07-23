package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/last-stone-weight-ii/
//------------------------Leetcode Problem 1049------------------------
func lastStoneWeightII(stones []int) int {
	var sum int
	for _, x := range stones {
		sum += x
	}
	half := sum >> 1
	dp := make([]bool, half+1)
	dp[0] = true
	for _, x := range stones {
		for i := half; i >= x; i-- {
			if dp[i-x] {
				dp[i] = true
			}
		}
	}
	for i := half; ; i-- {
		if dp[i] {
			return sum - (i << 1)
		}
	}
}

//------------------------Leetcode Problem 1049------------------------
/*
 * https://leetcode.cn/problems/last-stone-weight-ii/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2MB 在所有Go提交中击败了93.90%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		stones := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", lastStoneWeightII(stones))
	}
}
