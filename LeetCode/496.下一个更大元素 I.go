package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	var stack []int
	for i := len(nums2) - 1; i >= 0; i-- {
		x := nums2[i]
		for len(stack) > 0 && x >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			m[x] = stack[len(stack)-1]
		} else {
			m[x] = -1
		}
		stack = append(stack, x)
	}
	var ans []int
	for _, x := range nums1 {
		ans = append(ans, m[x])
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		nums1 := stringArrayToIntArray(strings.Fields(input.Text()))
		input.Scan()
		nums2 := stringArrayToIntArray(strings.Fields(input.Text()))
		ans := nextGreaterElement(nums1, nums2)
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
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：3.4MB 在所有Go提交中击败了25.31%的用户
**/
