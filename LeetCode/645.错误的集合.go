package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findErrorNums(nums []int) []int {
	mp := make([]int, len(nums)+1)
	for _, x := range nums {
		mp[x]++
	}
	ans := make([]int, 2)
	for i, x := range mp {
		if x > 1 {
			ans[0] = i
		} else if x == 0 {
			ans[1] = i
		}
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		ans := findErrorNums(stringArrayToIntArray(strings.Fields(input.Text())))
		for _, x := range ans {
			fmt.Printf("%d ", x)
		}
		fmt.Println()
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
 * 执行用时：32ms 在所有Go提交中击败了72.48%的用户
 * 占用内存：6.5MB 在所有Go提交中击败了53.69%的用户
**/
