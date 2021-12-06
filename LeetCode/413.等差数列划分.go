package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func numberOfArithmeticSlices(nums []int) int {
	if len(nums) <= 2 {
		return 0
	}
	var ans int
	var add int
	for i := 2; i < len(nums); i++ {
		if nums[i-1]-nums[i] == nums[i-2]-nums[i-1] {
			add++
			ans += add
		} else {
			add = 0
		}
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		nums := stringArrayToIntArray(strings.Fields(input.Text()))
		println(numberOfArithmeticSlices(nums))
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
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.2MB 在所有Go提交中击败了100.00%的用户
**/
