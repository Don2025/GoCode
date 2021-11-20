package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func findLHS(nums []int) int {
	var ans int
	cnt := make(map[int]int)
	for _, num := range nums {
		cnt[num]++
	}
	for num, c := range cnt {
		if cnt[num+1] > 0 {
			ans = max(ans, c+cnt[num+1])
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		println(findLHS(stringArrayToIntArray(strings.Fields(input.Text()))))
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
 * 执行用时：56ms 在所有Go提交中击败了60.98%的用户
 * 占用内存：7.2MB 在所有Go提交中击败了25.00%的用户
**/
