package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func maxSubsequence(nums []int, k int) []int {
	a := make([]int, len(nums))
	for i, x := range nums {
		a[i] = x
	}
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	var ans []int
	cnt := make(map[int]int, k)
	for i := 0; i < k; i++ {
		cnt[nums[i]]++
	}
	for _, x := range a {
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
		nums := stringArrayToIntArray(strings.Fields(input.Text()))
		input.Scan()
		k, _ := strconv.Atoi(input.Text())
		ans := maxSubsequence(nums, k)
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
