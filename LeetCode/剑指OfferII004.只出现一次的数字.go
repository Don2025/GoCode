package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func singleNumber(nums []int) int {
	m := map[int]int{}
	for _, x := range nums {
		m[x]++
	}
	for k, v := range m {
		if v == 1 {
			return k
		}
	}
	return 0
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		nums := stringArrayToIntArray(strings.Fields(input.Text()))
		println(singleNumber(nums))
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
 * 执行用时：4ms 在所有Go提交中击败了92.99%的用户
 * 占用内存：4MB 在所有Go提交中击败了14.81%的用户
**/
