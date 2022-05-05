package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func findKthLargest(nums []int, k int) int {
	n := len(nums)
	buildMaxHeap(nums, n)
	for i := len(nums) - 1; i > len(nums)-k; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		n--
		maxHepify(nums, 0, n)
	}
	return nums[0]
}

func buildMaxHeap(nums []int, n int) {
	for i := n >> 1; i >= 0; i-- {
		maxHepify(nums, i, n)
	}
}

func maxHepify(nums []int, i, n int) {
	l, r, largest := (i<<1)+1, (i+1)<<1, i
	if l < n && nums[l] > nums[largest] {
		largest = l
	}
	if r < n && nums[r] > nums[largest] {
		largest = r
	}
	if largest != i {
		nums[i], nums[largest] = nums[largest], nums[i]
		maxHepify(nums, largest, n)
	}
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		nums := stringArrayToIntArray(strings.Fields(input.Text()))
		input.Scan()
		k, _ := strconv.Atoi(input.Text())
		println(findKthLargest(nums, k))
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
 * 执行用时：4ms 在所有Go提交中击败了97.36%的用户
 * 占用内存：3.3MB 在所有Go提交中击败了88.36%的用户
**/
