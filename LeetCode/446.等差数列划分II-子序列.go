package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func numberOfArithmeticSlices(nums []int) int {
	var ans int
	mp := make([]map[int]int, len(nums))
	for i, x := range nums {
		mp[i] = map[int]int{}
		for j, y := range nums[:i] {
			diff := x - y
			cnt := mp[j][diff]
			ans += cnt
			mp[i][diff] += cnt + 1
		}
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		println(numberOfArithmeticSlices(stringArrayToIntArray(strings.Fields(input.Text()))))
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
 * 执行用时：176ms 在所有Go提交中击败了45.83%的用户
 * 占用内存：37.2MB 在所有Go提交中击败了58.33%的用户
**/
