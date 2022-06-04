package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func combinationSum(candidates []int, target int) [][]int {
	var ans [][]int
	var arr []int
	var dfs func(int, int)
	dfs = func(target, index int) {
		if index == len(candidates) {
			return
		}
		if target == 0 {
			ans = append(ans, append([]int(nil), arr...))
			return
		}
		dfs(target, index+1)
		if target-candidates[index] >= 0 {
			arr = append(arr, candidates[index])
			dfs(target-candidates[index], index)
			arr = arr[:len(arr)-1]
		}
	}
	dfs(target, 0)
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		candidates := stringArrayToIntArray(strings.Fields(input.Text()))
		input.Scan()
		target, _ := strconv.Atoi(input.Text())
		ans := combinationSum(candidates, target)
		for _, row := range ans {
			fmt.Printf("%v ", row)
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
 * 占用内存：2.6MB 在所有Go提交中击败了59.60%的用户
**/
