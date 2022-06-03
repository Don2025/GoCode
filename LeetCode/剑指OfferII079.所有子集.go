package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func subsets(nums []int) [][]int {
	var ans [][]int
	n := len(nums)
	for mask := 0; mask < 1<<n; mask++ {
		var a []int
		for i, x := range nums {
			if mask>>i&1 > 0 {
				a = append(a, x)
			}
		}
		ans = append(ans, append([]int(nil), a...))
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		nums := stringArrayToIntArray(strings.Fields(input.Text()))
		ans := subsets(nums)
		for _, row := range ans {
			fmt.Printf("%v ", row)
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
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.2MB 在所有Go提交中击败了9.78%的用户
**/
