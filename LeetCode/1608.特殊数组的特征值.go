package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func specialArray(nums []int) int {
	l, r := 1, 100
	for l <= r {
		mid := l + (r-l)>>1
		var cnt int
		for _, x := range nums {
			if x >= mid {
				cnt++
			}
		}
		if cnt == mid {
			return mid
		} else if cnt > mid {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		nums := stringArrayToIntArray(strings.Fields(input.Text()))
		println(specialArray(nums))
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
 * 占用内存：2MB 在所有Go提交中击败了89.34%的用户
**/
