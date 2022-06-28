package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"sort"
)

// https://leetcode.cn/problems/wiggle-sort-ii/
//------------------------Leetcode Problem 324------------------------
func wiggleSort(nums []int) {
	n := len(nums)
	arr := make([]int, n)
	copy(arr, nums)
	sort.Ints(arr)
	l, r := (n-1)>>1, n-1
	for i := 0; i < n; i++ {
		if i&1 == 0 {
			nums[i] = arr[l]
			l--
		} else {
			nums[i] = arr[r]
			r--
		}
	}
}

//------------------------Leetcode Problem 324------------------------
/*
 * https://leetcode.cn/problems/wiggle-sort-ii/
 * 执行用时：28ms 在所有Go提交中击败了82.33%的用户
 * 占用内存：6.7MB 在所有Go提交中击败了82.33%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		wiggleSort(nums)
		Printf("Output: %v\n", nums)
	}
}
