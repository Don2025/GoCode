package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"sort"
	"strconv"
)

// https://leetcode.cn/problems/find-subsequence-of-length-k-with-the-largest-sum/
// ------------------------Leetcode Problem 2099------------------------
func maxSubsequence(nums []int, k int) []int {
	a := make([]int, len(nums))
	for i, x := range nums {
		a[i] = x
	}
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	var ans []int
	cnt := make(map[int]int, k)
	for i := 0; i < k; i++ {
		cnt[nums[i]]++
	}
	for _, x := range a {
		if cnt[x] > 0 {
			ans = append(ans, x)
			cnt[x]--
		}
	}
	return ans
}

//------------------------Leetcode Problem 2099------------------------
/*
 * https://leetcode.cn/problems/find-subsequence-of-length-k-with-the-largest-sum/
 * 执行用时：4ms 在所有Go提交中击败了95.17%的用户
 * 占用内存：3.6MB 在所有Go提交中击败了39.08%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		scanner.Scan()
		k, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", maxSubsequence(nums, k))
	}
}
