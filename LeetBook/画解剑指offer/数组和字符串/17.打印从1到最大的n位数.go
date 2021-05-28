package LeetBook

import "math"

func printNumbers(n int) []int {
	n = int(math.Pow(10, float64(n)))
	a := make([]int, n-1)
	for i := 1; i < n; i++ {
		a[i-1] = i
	}
	return a
}

/*
 * 执行用时：8ms 在所有Go提交中击败了90.61%的用户
 * 占用内存：7MB 在所有Go提交中击败了87.71%的用户
**/
