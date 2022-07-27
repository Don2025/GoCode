package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"sort"
)

// https://leetcode.cn/problems/sort-integers-by-the-number-of-1-bits/
//------------------------Leetcode Problem 1356------------------------
func sortByBits(arr []int) []int {
	var count1 func(int) int
	count1 = func(n int) int {
		var cnt int
		for n != 0 {
			n = n & (n - 1)
			cnt++
		}
		return cnt
	}
	sort.Slice(arr, func(i, j int) bool {
		cnt1, cnt2 := count1(arr[i]), count1(arr[j])
		return cnt1 < cnt2 || cnt1 == cnt2 && arr[i] < arr[j]
	})
	return arr
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		arr := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", sortByBits(arr))
	}
}

//------------------------Leetcode Problem 1356------------------------
/*
 * https://leetcode.cn/problems/sort-integers-by-the-number-of-1-bits/
 * 执行用时：8ms 在所有Go提交中击败了60.81%的用户
 * 占用内存：3.4MB 在所有Go提交中击败了100.00%的用户
**/
