package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

//输入数字 n，按顺序打印出从 1 到最大的 n 位十进制数。比如输入 3，则打印出 1、2、3 一直到最大的 3 位数 999。
func printNumbers(n int) []int {
	n = int(math.Pow(10, float64(n)))
	a := make([]int, n-1)
	for i := 1; i < n; i++ {
		a[i-1] = i
	}
	return a
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		ans := printNumbers(n)
		for _, x := range ans {
			fmt.Printf("%d ", x)
		}
		println()
	}
}

/*
 * 执行用时：8ms 在所有Go提交中击败了90.61%的用户
 * 占用内存：7MB 在所有Go提交中击败了87.71%的用户
**/
