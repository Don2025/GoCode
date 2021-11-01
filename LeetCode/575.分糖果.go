package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func distributeCandies(candyType []int) int {
	mp := make(map[int]bool)
	n := len(candyType) >> 1
	for i := 0; i < len(candyType); i++ {
		mp[candyType[i]] = true
		if len(mp) >= n {
			return n
		}
	}
	return len(mp)
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		println(distributeCandies(stringArrayToIntArray(strings.Fields(input.Text()))))
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
 * 执行用时：140ms 在所有Go提交中击败了96.71%的用户
 * 占用内存：7.2MB 在所有Go提交中击败了75.66%的用户
**/
