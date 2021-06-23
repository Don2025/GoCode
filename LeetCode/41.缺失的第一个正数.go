package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func firstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); i++ {
		adjust(nums, i)
	}
	for i := 0; i < len(nums); i++ {
		if nums[i]-1 != i {
			return i + 1
		}
	}
	return len(nums) + 1
}

func adjust(nums []int, i int) {
	for nums[i] > 0 && nums[i] <= len(nums) && nums[i] != nums[nums[i]-1] {
		nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
	}
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		println(firstMissingPositive(stringArrayToIntArray(strings.Fields(input.Text()))))
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
 * 执行用时：112ms 在所有Go提交中击败了36.10%的用户
 * 占用内存：22MB 在所有Go提交中击败了10.59%的用户
**/
