package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func majorityElement(nums []int) []int {
	var ans []int
	cnt := map[int]int{}
	for _, x := range nums {
		cnt[x]++
	}
	for i, x := range cnt {
		if x > len(nums)/3 {
			ans = append(ans, i)
		}
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		ans := majorityElement(stringArrayToIntArray(strings.Fields(input.Text())))
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
 * 执行用时：8ms 在所有Go提交中击败了92.86%的用户
 * 占用内存：5MB 在所有Go提交中击败了23.63%的用户
**/
