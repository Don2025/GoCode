package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getRow(rowIndex int) []int {
	var ans []int
	num := 1
	for i := 0; i <= rowIndex; i++ {
		ans = append(ans, num)
		num = num * (rowIndex - i) / (i + 1)
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		rowIndex, _ := strconv.Atoi(input.Text())
		ans := getRow(rowIndex)
		for _, x := range ans {
			fmt.Printf("%d ", x)
		}
		fmt.Println()
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2MB 在所有Go提交中击败了64.20%的用户
**/
