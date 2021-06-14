package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func constructArr(a []int) []int {
	ans := make([]int, len(a))
	left := 1
	for i, x := range a {
		ans[i] = left
		left *= x
	}
	right := 1
	for i := len(a) - 1; i >= 0; i-- {
		ans[i] *= right
		right *= a[i]
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		arr := constructArr(stringArrayToIntArray(strings.Fields(input.Text())))
		for _, x := range arr {
			fmt.Printf("%d ", x)
		}
		println()
	}
}

/*
 * 执行用时：28ms 在所有Go提交中击败了81.77%的用户
 * 占用内存：9.2MB 在所有Go提交中击败了58.18%的用户
**/
