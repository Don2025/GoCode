package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"sort"
)

// https://leetcode.cn/problems/peak-index-in-a-mountain-array/
//------------------------Leetcode Problem 852------------------------
func peakIndexInMountainArray(arr []int) int {
	return sort.Search(len(arr)-1, func(i int) bool {
		return arr[i] > arr[i+1]
	})
}

//------------------------Leetcode Problem 852------------------------
/*
 * https://leetcode.cn/problems/peak-index-in-a-mountain-array/
 * 执行用时：16ms 在所有Go提交中击败了7.24%的用户
 * 占用内存：4.5MB 在所有Go提交中击败了63.82%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		arr := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", peakIndexInMountainArray(arr))
	}
}
