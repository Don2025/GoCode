package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func sortByBits(arr []int) []int {
	var count1 func(int) int
	count1 = func(n int) int {
		var cnt int
		for n != 0 {
			n = n & (n - 1)
			cnt++
		}
		return cnt
	}
	sort.Slice(arr, func(i, j int) bool {
		cnt1, cnt2 := count1(arr[i]), count1(arr[j])
		return cnt1 < cnt2 || cnt1 == cnt2 && arr[i] < arr[j]
	})
	return arr
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		ans := sortByBits(stringArrayToIntArray(strings.Fields(input.Text())))
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
 * 执行用时：8ms 在所有Go提交中击败了60.81%的用户
 * 占用内存：3.4MB 在所有Go提交中击败了100.00%的用户
**/
