package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func singleNumber(nums []int) []int {
	cnt := map[int]int{}
	for _, x := range nums {
		cnt[x]++
	}
	var ans []int
	for key, val := range cnt {
		if val == 1 {
			ans = append(ans, key)
		}
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		ans := singleNumber(stringArrayToIntArray(strings.Fields(input.Text())))
		for _, x := range ans {
			fmt.Printf("%d ", x)
		}
		fmt.Println()
	}
}

/*
 * 执行用时：8ms 在所有Go提交中击败了27.51%的用户
 * 占用内存：4.8MB 在所有Go提交中击败了17.03%的用户
**/
