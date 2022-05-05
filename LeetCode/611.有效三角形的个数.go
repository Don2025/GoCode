package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func triangleNumber(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	var ans int
	for i := range nums {
		k := i
		for j := i + 1; j < n; j++ {
			for k+1 < n && nums[k+1] < nums[i]+nums[j] {
				k++
			}
			ans += max(0, k-j)
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		nums := stringArrayToIntArray(strings.Fields(input.Text()))
		println(triangleNumber(nums))
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
 * 执行用时：40ms 在所有Go提交中击败了56.51%的用户
 * 占用内存：3MB 在所有Go提交中击败了53.90%的用户
**/
