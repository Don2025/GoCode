package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func subarraySum(nums []int, k int) int {
	var ans int
	for i := 0; i < len(nums); i++ {
		var sum int
		for j := i; j >= 0; j-- {
			sum += nums[j]
			if sum == k {
				ans++
			}
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
		println(subarraySum(nums, k))
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
 * 执行用时：1096ms 在所有Go提交中击败了6.60%的用户
 * 占用内存：6.5MB 在所有Go提交中击败了80.97%的用户
**/
