package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func lexicalOrder(n int) []int {
	ans := make([]int, n)
	num := 1
	for i := range ans {
		ans[i] = num
		if num*10 <= n {
			num *= 10
		} else {
			for num%10 == 9 || num+1 > n {
				num /= 10
			}
			num++
		}
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		ans := lexicalOrder(n)
		fmt.Printf("%v\n", ans)
	}
}

/*
 * 执行用时：4ms 在所有Go提交中击败了99.09%的用户
 * 占用内存：6.3MB 在所有Go提交中击败了100.00%的用户
**/
