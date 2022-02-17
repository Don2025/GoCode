package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func lengthOfLIS(nums []int) int {
	var dp []int
	for _, num := range nums {
		if len(dp) == 0 || dp[len(dp)-1] < num {
			dp = append(dp, num)
		} else {
			l, r := 0, len(dp)-1
			pos := r
			for l <= r {
				mid := (l + r) >> 1
				if dp[mid] >= num {
					pos = mid
					r = mid - 1
				} else {
					l = mid + 1
				}
			}
			dp[pos] = num
		}
	}
	return len(dp)
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		nums := stringArrayToIntArray(strings.Fields(input.Text()))
		println(lengthOfLIS(nums))
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
 * 执行用时：4ms 在所有Go提交中击败了98.99%的用户
 * 占用内存：3.4MB 在所有Go提交中击败了87.92%的用户
**/
