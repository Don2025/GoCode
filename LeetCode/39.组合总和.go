package main

import (
	"bufio"
	"fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
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
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		candidates := utils.StringToInts(scanner.Text())
		scanner.Scan()
		target, _ := strconv.Atoi(scanner.Text())
		ans := combinationSum(candidates, target)
		for _, x := range ans {
			fmt.Printf("%v ", x)
		}
		fmt.Println()
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.6MB 在所有Go提交中击败了73.88%的用户
**/
