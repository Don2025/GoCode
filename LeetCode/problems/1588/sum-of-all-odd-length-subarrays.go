package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/sum-of-all-odd-length-subarrays/
//------------------------Leetcode Problem 1588------------------------
func sumOddLengthSubarrays(arr []int) int {
	n := len(arr)
	prefix := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] + arr[i]
	}
	var ans int
	for i := 0; i < n; i++ {
		for j := i; j >= 0; j -= 2 {
			ans += prefix[i+1] - prefix[j]
		}
	}
	return ans
}

//------------------------Leetcode Problem 1588------------------------
/*
 * https://leetcode.cn/problems/sum-of-all-odd-length-subarrays/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.1MB 在所有Go提交中击败了43.59%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		arr := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", sumOddLengthSubarrays(arr))
	}
}
