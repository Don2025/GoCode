package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func findMaxLength(nums []int) int {
	var cnt, ans int
	mp := map[int]int{0: -1}
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			cnt++
		} else {
			cnt--
		}
		if preidx, has := mp[cnt]; has {
			t := i - preidx
			if t > ans {
				ans = t
			}
		} else {
			mp[cnt] = i
		}
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		a := stringArrayToIntArray(strings.Fields(input.Text()))
		println(findMaxLength(a))
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
 * 执行用时：132ms 在所有Go提交中击败了22.22%的用户
 * 占用内存：7.5MB 在所有Go提交中击败了27.78%的用户
**/
