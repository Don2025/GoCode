package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

/* 把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。
 * 输入一个递增排序的数组的一个旋转，输出旋转数组的最小元素。
**/
func minArray(numbers []int) int {
	left, right := 0, len(numbers)-1
	for left < right {
		mid := left + (right-left)/2
		if numbers[mid] < numbers[right] {
			right = mid
		} else if numbers[mid] > numbers[right] {
			left = mid + 1
		} else {
			right--
		}
	}
	return numbers[left]
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		a := stringArrayToIntArray(strings.Fields(input.Text()))
		println(minArray(a))
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
 * 执行用时：4ms 在所有Go提交中击败了91.75%的用户
 * 占用内存：3.1MB 在所有Go提交中击败了100.00%的用户
**/
