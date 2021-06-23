package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// grayCode 格雷编码的生成过程: G(i)=i^(i/2)
func grayCode(n int) []int {
	var ans []int
	for i := 0; i < (1 << n); i++ {
		ans = append(ans, i^i>>1)
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		ans := grayCode(n)
		for _, x := range ans {
			fmt.Printf("%d ", x)
		}
		println()
	}
}

/*
 * 执行用时：8ms 在所有Go提交中击败了98.16%的用户
 * 占用内存：6.7MB 在所有Go提交中击败了73.80%的用户
**/
