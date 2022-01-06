package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func jump(nums []int) int {
	var ans, end, maxDis int
	for i := 0; i < len(nums)-1; i++ {
		maxDis = max(maxDis, i+nums[i])
		if i == end {
			end = maxDis
			ans++
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
		println(jump(nums))
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
 * 执行用时：8ms 在所有Go提交中击败了99.39%的用户
 * 占用内存：5.5MB 在所有Go提交中击败了100.00%的用户
**/
