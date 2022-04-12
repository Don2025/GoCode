package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var ans [][]int
	var path []int
	var dfs func(int, int, int)
	dfs = func(index, curSum, target int) {
		if target == curSum {
			ans = append(ans, append([]int(nil), path...))
			return
		}
		for i := index; i < len(candidates); i++ {
			if i != index && candidates[i] == candidates[i-1] {
				continue
			}
			curSum += candidates[i]
			if curSum > target {
				break
			}
			path = append(path, candidates[i])
			dfs(i+1, curSum, target)
			curSum -= candidates[i]
			path = path[:len(path)-1]
		}
		return
	}
	dfs(0, 0, target)
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		candidates := stringArrayToIntArray(strings.Fields(input.Text()))
		input.Scan()
		target, _ := strconv.Atoi(input.Text())
		ans := combinationSum2(candidates, target)
		for _, x := range ans {
			fmt.Printf("%v ", x)
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
 * 占用内存：2.4MB 在所有Go提交中击败了87.68%的用户
**/
