package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, x := range nums {
		m[x]++
	}
	var ans []int
	for x, _ := range m {
		ans = append(ans, x)
	}
	sort.Slice(ans, func(i, j int) bool {
		return m[ans[i]] > m[ans[j]]
	})
	return ans[:k]
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		nums := stringArrayToIntArray(strings.Fields(input.Text()))
		input.Scan()
		k, _ := strconv.Atoi(input.Text())
		fmt.Printf("%v\n", topKFrequent(nums, k))
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
 * 执行用时：12ms 在所有Go提交中击败了82.90%的用户
 * 占用内存：5.3MB 在所有Go提交中击败了86.96%的用户
**/
