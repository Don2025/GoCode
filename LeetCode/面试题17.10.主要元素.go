package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

func majorityElement(nums []int) int {
	key := nums[0]
	var cnt int
	for _, x := range nums {
		if key == x {
			cnt++
		} else {
			cnt--
		}
		if cnt < 0 {
			key = x
			cnt = 0
		}
	}
	cnt = 0
	for _, x := range nums {
		if x == key {
			cnt++
		}
	}
	if cnt > len(nums)/2 {
		return key
	}
	return -1
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		println(majorityElement(stringArrayToIntArray(strings.Fields(input.Text()))))
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
 * 执行用时：16ms 在所有Go提交中击败了97.56%的用户
 * 占用内存：6MB 在所有Go提交中击败了78.86%的用户
**/
