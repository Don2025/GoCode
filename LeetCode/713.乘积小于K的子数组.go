package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func numSubarrayProductLessThanK(nums []int, k int) int {
	if k == 0 || k == 1 {
		return 0
	}
	prod, ans := 1, 0 // prod存储nums[l]~nums[r]的累积
	for l, r := 0, 0; r < len(nums); r++ {
		prod *= nums[r]
		for prod >= k {
			prod /= nums[l]
			l++
		}
		ans += r - l + 1
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		nums := stringArrayToIntArray(strings.Fields(input.Text()))
		input.Scan()
		k, _ := strconv.Atoi(input.Text())
		println(numSubarrayProductLessThanK(nums, k))
	}
}

/*
 * 执行用时：76ms 在所有Go提交中击败了47.90%的用户
 * 占用内存：7.2MB 在所有Go提交中击败了37.89%的用户
**/
