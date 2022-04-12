package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func subsets(nums []int) [][]int {
	n := len(nums)
	var ans [][]int
	for i := 0; i < 1<<n; i++ {
		var arr []int
		for j := range nums {
			if i>>j&1 > 0 {
				arr = append(arr, nums[j])
			}
		}
		ans = append(ans, append([]int(nil), arr...))
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		nums := stringArrayToIntArray(strings.Fields(input.Text()))
		ans := subsets(nums)
		for _, x := range ans {
			fmt.Printf("%v ", x)
		}
		fmt.Println()
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
 * 占用内存：2.2MB 在所有Go提交中击败了25.72%的用户
**/
