package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func intersect(nums1 []int, nums2 []int) []int {
	cnt := make(map[int]int)
	var ans []int
	for _, x := range nums1 {
		cnt[x]++
	}
	for _, x := range nums2 {
		if cnt[x] > 0 {
			ans = append(ans, x)
			cnt[x]--
		}
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		nums1 := stringArrayToIntArray(strings.Fields(input.Text()))
		input.Scan()
		nums2 := stringArrayToIntArray(strings.Fields(input.Text()))
		ans := intersect(nums1, nums2)
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
 * 执行用时：4ms 在所有Go提交中击败了73.88%的用户
 * 占用内存：3.1MB 在所有Go提交中击败了36.79%的用户
**/
