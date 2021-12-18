package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param A int整型一维数组
 * @param n int整型
 * @return int整型
 */
func getDis(A []int, n int) int {
	var ans int
	minVal := A[0]
	for i := 1; i < len(A); i++ {
		ans = max(ans, A[i]-minVal)
		minVal = min(minVal, A[i])
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		nums := stringArrayToIntArray(strings.Fields(input.Text()))
		input.Scan()
		n, _ := strconv.Atoi(input.Text())
		println(getDis(nums, n))
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
 * 运行时间：74ms 超过100.00%用Go提交的代码
 * 占用内存：20292KB 超过0.00%用Go提交的代码
**/
