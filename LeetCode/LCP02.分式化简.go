package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func fraction(cont []int) []int {
	up := 1
	down := cont[len(cont)-1]
	for i := len(cont) - 2; i >= 0; i-- {
		up = cont[i]*down + up
		up, down = down, up
	}
	return []int{down, up}
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		ans := fraction(stringArrayToIntArray(strings.Fields(input.Text())))
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
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了68.75%的用户
**/
