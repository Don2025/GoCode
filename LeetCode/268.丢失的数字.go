package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func missingNumber(nums []int) int {
	hash := make([]bool, len(nums)+1)
	for _, x := range nums {
		hash[x] = true
	}
	var ans int
	for i, x := range hash {
		if x == false {
			ans = i
		}
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		println(missingNumber(stringArrayToIntArray(strings.Fields(input.Text()))))
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
 * 执行用时：16ms 在所有Go提交中击败了86.81%的用户
 * 占用内存：6.2MB 在所有Go提交中击败了11.90%的用户
**/
