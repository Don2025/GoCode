package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)>>1
		if target == nums[mid] {
			return mid
		} else if target < nums[mid] {
			r = mid - 1
		} else { // target > nums[mid]
			l = mid + 1
		}
	}
	return l
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		nums := stringArrayToIntArray(strings.Fields(input.Text()))
		input.Scan()
		target, _ := strconv.Atoi(input.Text())
		println(searchInsert(nums, target))
	}
}

func stringArrayToIntArray(strings []string) []int {
	var arr []int
	for _, x := range strings {
		n, _ := strconv.Atoi(x)
		arr = append(arr, n)
	}
	/*
	 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
	 * 占用内存：2MB 在所有Go提交中击败了56.24%的用户
	**/

	return arr
}

/*
 * 执行用时：4ms 在所有Go提交中击败了71.71%的用户
 * 占用内存：3MB 在所有Go提交中击败了49.57%的用户
**/
