package main

import "fmt"

func getSum(a int, b int) int {
	return a + b
}

func main() {
	var a, b int
	for {
		fmt.Scanf("%d %d", &a, &b)
		println(getSum(a, b))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了65.41%的用户
**/
