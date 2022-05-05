package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func purchasePlans(nums []int, target int) int {
	const mod = 1e9 + 7
	sort.Ints(nums)
	l, r := 0, len(nums)-1
	var sum int
	for l != r {
		if nums[l]+nums[r] > target {
			r--
		} else {
			sum += r - l
			l++
		}
	}
	return sum % mod
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		nums := stringArrayToIntArray(strings.Fields(input.Text()))
		input.Scan()
		target, _ := strconv.Atoi(input.Text())
		println(purchasePlans(nums, target))
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
 * 执行用时：136ms 在所有Go提交中击败了77.27%的用户
 * 占用内存：7.8MB 在所有Go提交中击败了78.79%的用户
**/
