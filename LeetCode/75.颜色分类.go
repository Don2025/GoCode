package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func sortColors(nums []int) {
	i, j := -1, -1
	for x := range nums {
		if nums[x] < 2 {
			j++
			nums[x], nums[j] = nums[j], nums[x]
			if nums[j] < 1 {
				i++
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		nums := stringArrayToIntArray(strings.Fields(input.Text()))
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
 * 占用内存：1.9MB 在所有Go提交中击败了100.00%的用户
**/
