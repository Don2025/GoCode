package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			if nums[left] != target {
				left++
			} else if nums[right] != target {
				right--
			} else {
				break
			}
		}
	}
	return right - left + 1
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		nums := stringArrayToIntArray(strings.Fields(input.Text()))
		input.Scan()
		target, _ := strconv.Atoi(input.Text())
		println(search(nums, target))
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
 * 执行用时：8ms 在所有Go提交中击败了89.40%的用户
 * 占用内存：3.9MB 在所有Go提交中击败了60.12%的用户
**/
