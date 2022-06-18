package main

import (
	"bufio"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
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
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		scanner.Scan()
		k, _ := strconv.Atoi(scanner.Text())
		println(findKthLargest(nums, k))
	}
}

/*
 * 执行用时：4ms 在所有Go提交中击败了97.36%的用户
 * 占用内存：3.3MB 在所有Go提交中击败了88.36%的用户
**/
