package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/* 输入一个递增排序的数组和一个数字s，在数组中查找两个数，使得它们的和正好是s。
 * 如果有多对数字的和等于s，则输出任意一对即可。
**/
func twoSum(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		sum := nums[left] + nums[right]
		if sum < target {
			left++
		} else if sum > target {
			right--
		} else {
			return []int{nums[left], nums[right]}
		}
	}
	return []int{0}
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		nums := stringArrayToIntArray(strings.Fields(input.Text()))
		input.Scan()
		target, _ := strconv.Atoi(input.Text())
		ans := twoSum(nums, target)
		for _, x := range ans {
			fmt.Printf("%d ", x)
		}
		println()
	}
}

func stringArrayToIntArray(strings []string) []int {
	var arr []int
	for _, x := range strings {
		n, _ := strconv.Atoi(x)
		arr = append(arr, n)
	}
	return arr
}

/*
 * 执行用时：220ms 在所有Go提交中击败了66.13%的用户
 * 占用内存：8.6MB 在所有Go提交中击败了66.99%的用户
**/
