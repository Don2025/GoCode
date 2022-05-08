package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findDuplicates(nums []int) []int {
	for i := range nums {
		for nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	var ans []int
	for i := range nums {
		if nums[i]-1 != i {
			ans = append(ans, nums[i])
		}
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		nums := findDuplicates(stringArrayToIntArray(strings.Fields(input.Text())))
		fmt.Printf("%v\n", nums)
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
 * 执行用时：44ms 在所有Go提交中击败了50.88%的用户
 * 占用内存：7.5MB 在所有Go提交中击败了63.16%的用户
**/
