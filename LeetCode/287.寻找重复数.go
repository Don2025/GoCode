package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func findDuplicate(nums []int) int {
	var slow, fast int
	for slow, fast = nums[slow], nums[nums[fast]]; slow != fast; slow, fast = nums[slow], nums[nums[fast]] {
	}
	for slow = 0; slow != fast; {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		nums := stringArrayToIntArray(strings.Fields(input.Text()))
		println(findDuplicate(nums))
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
 * 执行用时：76ms 在所有Go提交中击败了92.51%的用户
 * 占用内存：8.2MB 在所有Go提交中击败了49.62%的用户
**/
