package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func findContinuousSequence(target int) [][]int {
	var ans [][]int
	var sum int
	for l, r := 1, 1; r < target; r++ {
		sum += r
		for sum > target {
			sum -= l
			l++
		}
		if sum == target {
			arr := make([]int, r-l+1)
			for i := range arr {
				arr[i] = l + i
			}
			ans = append(ans, arr)
		}
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		ans := findContinuousSequence(n)
		for _, x := range ans {
			fmt.Printf("%v ", x)
		}
		fmt.Println()
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2MB 在所有Go提交中击败了90.62%的用户
**/
