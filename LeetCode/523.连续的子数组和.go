package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func checkSubarraySum(nums []int, k int) bool {
	mp := map[int]int{0: -1}
	var mod int
	for i, x := range nums {
		mod = (mod + x) % k
		if preidx, has := mp[mod]; has {
			if i-preidx >= 2 {
				return true
			}
		} else {
			mp[mod] = i
		}
	}
	return false
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		a := stringArrayToIntArray(strings.Fields(input.Text()))
		input.Scan()
		n, _ := strconv.Atoi(input.Text())
		println(checkSubarraySum(a, n))
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
 * 执行用时：228ms 在所有Go提交中击败了5.04%的用户
 * 占用内存：9.2MB 在所有Go提交中击败了38.13%的用户
**/
