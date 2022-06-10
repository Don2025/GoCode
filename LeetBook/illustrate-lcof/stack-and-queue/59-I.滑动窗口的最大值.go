package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func maxSlidingWindow(nums []int, k int) []int {
	var ans []int
	if len(nums) == 0 {
		return ans
	}
	for i := 0; i < len(nums)-k+1; i++ {
		ans = append(ans, getMaxValue(nums[i:i+k]))
	}
	return ans
}

func getMaxValue(arr []int) int {
	max := -math.MaxInt32
	for _, x := range arr {
		if max < x {
			max = x
		}
	}
	return max
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		nums := stringArrayToIntArray(strings.Fields(input.Text()))
		input.Scan()
		k, _ := strconv.Atoi(input.Text())
		ans := maxSlidingWindow(nums, k)
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
 * 执行用时：40ms 在所有Go提交中击败了18.79%的用户
 * 占用内存：6.3MB 在所有Go提交中击败了87.29%的用户
**/
