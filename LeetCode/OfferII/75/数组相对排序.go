package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/0H97ZC/
// ------------------------剑指 Offer II Problem 75------------------------
func relativeSortArray(arr1 []int, arr2 []int) []int {
	upper := 0
	for _, v := range arr1 {
		if v > upper {
			upper = v
		}
	}
	frequency := make([]int, upper+1)
	for _, v := range arr1 {
		frequency[v]++
	}

	ans := make([]int, 0, len(arr1))
	for _, v := range arr2 {
		for ; frequency[v] > 0; frequency[v]-- {
			ans = append(ans, v)
		}
	}
	for v, freq := range frequency {
		for ; freq > 0; freq-- {
			ans = append(ans, v)
		}
	}
	return ans
}

// ------------------------剑指 Offer II Problem 75------------------------
/*
 * https://leetcode.cn/problems/0H97ZC/
 * 执行用时： 0ms 在所有Go提交中击败了 100.00%的用户
 * 占用内存： 2.2MB 在所有Go提交中击败了 45.83%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	Printf("Input a line of numbers separated by space:")
	for scanner.Scan() {
		arr1 := utils.StringToInts(scanner.Text())
		Printf("Input a line of numbers separated by space:")
		scanner.Scan()
		arr2 := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", relativeSortArray(arr1, arr2))
		Printf("Input a line of numbers separated by space:")
	}
}
