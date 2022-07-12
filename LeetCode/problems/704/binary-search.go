package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/binary-search/
//------------------------Leetcode Problem 704------------------------
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)>>1
		if target < nums[mid] {
			r = mid - 1
		} else if target > nums[mid] {
			l = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

//------------------------Leetcode Problem 704------------------------
/*
 * https://leetcode.cn/problems/binary-search/
 * 执行用时：36ms 在所有Go提交中击败了39.64%的用户
 * 占用内存：6.7MB 在所有Go提交中击败了44.15%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		scanner.Scan()
		target, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", search(nums, target))
	}
}
