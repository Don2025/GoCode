package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func productExceptSelf(nums []int) []int {
	ans := make([]int, len(nums))
	for i := range ans {
		ans[i] = 1
	}
	left, right := 1, 1
	for i := 0; i < len(nums); i++ {
		ans[i] *= left
		left *= nums[i]
		ans[len(nums)-i-1] *= right
		right *= nums[len(nums)-i-1]
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		ans := productExceptSelf(stringArrayToIntArray(strings.Fields(input.Text())))
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
 * 执行用时：32ms 在所有Go提交中击败了81.68%的用户
 * 占用内存：7.6MB 在所有Go提交中击败了52.90%的用户
**/
