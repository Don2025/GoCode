package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	ans := make([]int, n)
	var stack []int
	for i := 0; i < n; i++ {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			preIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[preIndex] = i - preIndex
		}
		stack = append(stack, i)
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		temperatures := stringArrayToIntArray(strings.Fields(input.Text()))
		ans := dailyTemperatures(temperatures)
		fmt.Printf("%v\n", ans)
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
 * 执行用时：108ms 在所有Go提交中击败了88.00%的用户
 * 占用内存：8.6MB 在所有Go提交中击败了84.80%的用户
**/
