package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func findKthPositive(arr []int, k int) int {
	if arr[0] > k {
		return k
	}
	l, r := 0, len(arr)
	for l < r {
		mid := l + (r-l)>>1
		if arr[mid]-mid-1 < k {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return k + l
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		arr := stringArrayToIntArray(strings.Fields(input.Text()))
		input.Scan()
		k, _ := strconv.Atoi(input.Text())
		println(findKthPositive(arr, k))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.7MB 在所有Go提交中击败了13.79%的用户
**/
