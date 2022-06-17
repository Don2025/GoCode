package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"sort"
	"strconv"
)

// https://leetcode.cn/problems/find-k-th-smallest-pair-distance/
//------------------------Leetcode Problem 719------------------------
func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)
	return sort.Search(nums[len(nums)-1]-nums[0], func(mid int) bool {
		cnt := 0
		for j, num := range nums {
			i := sort.SearchInts(nums[:j], num-mid)
			cnt += j - i
		}
		return cnt >= k
	})
}

//------------------------Leetcode Problem 719------------------------
/*
 * https://leetcode.cn/problems/find-k-th-smallest-pair-distance/
 * 执行用时：16ms 在所有Go提交中击败了47.21%的用户
 * 占用内存：3.2MB 在所有Go提交中击败了67.77%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	Printf("Input a line of numbers separated by space:")
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		Printf("Input a number:")
		scanner.Scan()
		k, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", smallestDistancePair(nums, k))
		Printf("Input a line of numbers separated by space:")
	}
}
