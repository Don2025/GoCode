package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/special-array-with-x-elements-greater-than-or-equal-x/
//------------------------Leetcode Problem 1608------------------------
func specialArray(nums []int) int {
	l, r := 1, 100
	for l <= r {
		mid := l + (r-l)>>1
		var cnt int
		for _, x := range nums {
			if x >= mid {
				cnt++
			}
		}
		if cnt == mid {
			return mid
		} else if cnt > mid {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}

//------------------------Leetcode Problem 1608------------------------
/*
 * https://leetcode.cn/problems/special-array-with-x-elements-greater-than-or-equal-x/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2MB 在所有Go提交中击败了89.34%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", specialArray(nums))
	}
}
