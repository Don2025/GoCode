package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func singleNumbers(nums []int) []int {
	m := make(map[int]int, len(nums))
	for _, x := range nums {
		if _, has := m[x]; !has {
			m[x] = x
		} else {
			delete(m, x)
		}
	}
	var ans []int
	for x := range m {
		ans = append(ans, x)
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		ans := singleNumbers(stringArrayToIntArray(strings.Fields(input.Text())))
		for _, x := range ans {
			fmt.Printf("%d ", x)
		}
		println()
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
 * 执行用时：28ms 在所有Go提交中击败了19.39%的用户
 * 占用内存：6.5MB 在所有Go提交中击败了13.46%的用户
**/
