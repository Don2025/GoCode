package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func smallestRangeI(nums []int, k int) int {
	minVal, maxVal := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < minVal {
			minVal = nums[i]
		} else if nums[i] > maxVal {
			maxVal = nums[i]
		}
	}
	ans := maxVal - minVal - (k << 1)
	if ans > 0 {
		return ans
	}
	return 0
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		nums := stringArrayToIntArray(strings.Fields(input.Text()))
		input.Scan()
		k, _ := strconv.Atoi(input.Text())
		println(smallestRangeI(nums, k))
	}
}

/*
 * 执行用时：12ms 在所有Go提交中击败了84.21%的用户
 * 占用内存：5.9MB 在所有Go提交中击败了92.98%的用户
**/
