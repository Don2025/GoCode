package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	ans := make([][]int, 0)
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		k := n - 1
		target := -1 * nums[i]
		for j := i + 1; j < n; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			for j < k && nums[j]+nums[k] > target {
				k--
			}
			if j == k {
				break
			}
			if nums[j]+nums[k] == target {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
			}
		}
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		nums := threeSum(stringArrayToIntArray(strings.Fields(input.Text())))
		for _, row := range nums {
			fmt.Printf("%v ", row)
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
 * 执行用时：24ms 在所有Go提交中击败了97.58%的用户
 * 占用内存：7.3MB 在所有Go提交中击败了58.67%的用户
**/
