package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	m--
	n--
	for i := len(nums1) - 1; m >= 0 && n >= 0; i-- {
		if nums1[m] > nums2[n] {
			nums1[i] = nums1[m]
			m--
		} else {
			nums1[i] = nums2[n]
			n--
		}
	}
	if n >= 0 {
		copy(nums1[0:n+1], nums2[0:n+1])
	}
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		nums1 := stringArrayToIntArray(strings.Fields(input.Text()))
		input.Scan()
		m, _ := strconv.Atoi(input.Text())
		input.Scan()
		nums2 := stringArrayToIntArray(strings.Fields(input.Text()))
		input.Scan()
		n, _ := strconv.Atoi(input.Text())
		merge(nums1, m, nums2, n)
		for _, x := range nums1 {
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
 * 占用内存：2.3MB 在所有Go提交中击败了74.26%的用户
**/
