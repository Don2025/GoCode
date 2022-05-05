package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func isStraight(nums []int) bool {
	var joker int
	sort.Ints(nums)
	for i := 0; i < 4; i++ {
		if nums[i] == 0 {
			joker++ //大小王
		} else if nums[i] == nums[i+1] {
			return false
		}
	}
	return nums[4]-nums[joker] < 5
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		nums := stringArrayToIntArray(strings.Fields(input.Text()))
		println(isStraight(nums))
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
 * 占用内存：1.8MB 在所有Go提交中击败了75.21%的用户
**/
