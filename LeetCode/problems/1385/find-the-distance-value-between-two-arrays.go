package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"sort"
	"strconv"
)

// https://leetcode.cn/problems/find-the-distance-value-between-two-arrays/
//------------------------Leetcode Problem 1385------------------------
func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	if d < 0 {
		d = -d
	}
	sort.Ints(arr1)
	sort.Ints(arr2)
	var ans, l, r int
	for i := range arr1 {
		l = l + sort.Search(len(arr2[l:]), func(index int) bool {
			return arr2[l+index] >= arr1[i]-d
		})
		r = r + sort.Search(len(arr2[r:]), func(index int) bool {
			return arr2[r+index] > arr1[i]+d
		})
		if l >= r {
			ans++
		}
	}
	return ans
}

//------------------------Leetcode Problem 1385------------------------
/*
 * https://leetcode.cn/problems/find-the-distance-value-between-two-arrays/
 * 执行用时：8ms 在所有Go提交中击败了54.81%的用户
 * 占用内存：3.6MB 在所有Go提交中击败了12.50%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		arr1 := utils.StringToInts(scanner.Text())
		scanner.Scan()
		arr2 := utils.StringToInts(scanner.Text())
		scanner.Scan()
		d, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", findTheDistanceValue(arr1, arr2, d))
	}
}
