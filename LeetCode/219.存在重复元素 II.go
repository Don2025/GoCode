package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func containsNearbyDuplicate(nums []int, k int) bool {
	dict := make(map[int]int)
	for i, num := range nums {
		if x, ok := dict[num]; ok && i-x <= k {
			return true
		}
		dict[num] = i
	}
	return false
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		nums := stringArrayToIntArray(strings.Fields(input.Text()))
		input.Scan()
		k, _ := strconv.Atoi(input.Text())
		println(containsNearbyDuplicate(nums, k))
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
 * 执行用时：104ms 在所有Go提交中击败了73.19%的用户
 * 占用内存：12MB 在所有Go提交中击败了20.06%的用户
**/
