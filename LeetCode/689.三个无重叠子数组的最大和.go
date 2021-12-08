package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func maxSumOfThreeSubarrays(nums []int, k int) []int {
	var ans []int
	var sum1, maxSum1, maxSum1Idx int                  //第一个滑动窗口
	var sum2, maxSum12, maxSum12Idx1, maxSum12Idx2 int //第二个滑动窗口
	var sum3, maxSum123 int
	for i := k * 2; i < len(nums); i++ {
		sum1 += nums[i-k*2]
		sum2 += nums[i-k]
		sum3 += nums[i]
		if i >= k*3-1 {
			if sum1 > maxSum1 {
				maxSum1 = sum1
				maxSum1Idx = i - k*3 + 1
			}
			if maxSum1+sum2 > maxSum12 {
				maxSum12 = maxSum1 + sum2
				maxSum12Idx1, maxSum12Idx2 = maxSum1Idx, i-k*2+1
			}
			if maxSum12+sum3 > maxSum123 {
				maxSum123 = maxSum12 + sum3
				ans = []int{maxSum12Idx1, maxSum12Idx2, i - k + 1}
			}
			sum1 -= nums[i-k*3+1]
			sum2 -= nums[i-k*2+1]
			sum3 -= nums[i-k+1]
		}
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		nums := stringArrayToIntArray(strings.Fields(input.Text()))
		input.Scan()
		k, _ := strconv.Atoi(input.Text())
		ans := maxSumOfThreeSubarrays(nums, k)
		for _, x := range ans {
			fmt.Printf("%d ", x)
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
 * 执行用时：24ms 在所有Go提交中击败了79.31%的用户
 * 占用内存：6.2MB 在所有Go提交中击败了100.00%的用户
**/
