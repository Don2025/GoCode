package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sortedSquares(nums []int) []int {
	ans := make([]int, len(nums))
	l, r := 0, len(nums)-1
	for i := len(nums) - 1; l <= r; i-- {
		if abs(nums[l]) > abs(nums[r]) {
			ans[i] = pow2(nums[l])
			l++
		} else {
			ans[i] = pow2(nums[r])
			r--
		}
	}
	return ans
}

func abs(n int) int {
	if n < 0 {
		n = -n
	}
	return n
}

func pow2(n int) int {
	return n * n
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		nums := stringArrayToIntArray(strings.Fields(input.Text()))
		ans := sortedSquares(nums)
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
 * 执行用时：28ms 在所有Go提交中击败了80.78%的用户
 * 占用内存：6.8MB 在所有Go提交中击败了64.13%的用户
**/
