package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func countQuadruplets(nums []int) int {
	cnt := make(map[int]int)
	var ans int
	for i := len(nums) - 3; i > 0; i-- {
		for _, x := range nums[i+2:] {
			cnt[x-nums[i+1]]++
		}
		for _, x := range nums[:i] {
			if sum := x + nums[i]; cnt[sum] > 0 {
				ans += cnt[sum]
			}
		}
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		nums := stringArrayToIntArray(strings.Fields(input.Text()))
		println(countQuadruplets(nums))
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
 * 占用内存：2MB 在所有Go提交中击败了56.24%的用户
**/
