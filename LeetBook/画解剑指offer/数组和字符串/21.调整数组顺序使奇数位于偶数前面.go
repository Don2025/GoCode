package main

import "fmt"

//输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有奇数位于数组的前半部分，所有偶数位于数组的后半部分。
func exchange(nums []int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		for left < right && (nums[left]&1 == 1) {
			left++
		}
		for left < right && (nums[right]&1 == 0) {
			right--
		}
		nums[left], nums[right] = nums[right], nums[left]
	}
	return nums
}

func main() {
	nums := []int{1, 2, 3, 4}
	ans := exchange(nums)
	for _, x := range ans {
		fmt.Printf("%d ", x)
	}
}

/*
 * 执行用时：28ms 在所有Go提交中击败了44.12%的用户
 * 占用内存：6.8MB 在所有Go提交中击败了54.19%的用户
**/
