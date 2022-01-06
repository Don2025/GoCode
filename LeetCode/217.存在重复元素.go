package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func containsDuplicate(nums []int) bool {
	visited := make(map[int]bool)
	for _, num := range nums {
		if visited[num] {
			return true
		}
		visited[num] = true
	}
	return false
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		nums := stringArrayToIntArray(strings.Fields(input.Text()))
		println(containsDuplicate(nums))
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
 * 执行用时：92ms 在所有Go提交中击败了17.05%的用户
 * 占用内存：8.8MB 在所有Go提交中击败了15.83%的用户
**/
