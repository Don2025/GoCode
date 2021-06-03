package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func totalHammingDistance(nums []int) int {
	ans, n := 0, len(nums)
	for i := 0; i < 30; i++ {
		t := 0
		for _, val := range nums {
			t += val >> i & 1
		}
		ans += t * (n - t)
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		a := stringArrayToIntArray(strings.Fields(input.Text()))
		println(totalHammingDistance(a))
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
 * 执行用时：36ms 在所有Go提交中击败了68.09%的用户
 * 占用内存：6.5MB 在所有Go提交中击败了53.19%的用户
**/
