package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"sort"
	"strconv"
)

// https://leetcode.cn/problems/combination-sum-ii/
//------------------------Leetcode Problem 40------------------------
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
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		candidates := utils.StringToInts(scanner.Text())
		scanner.Scan()
		target, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", combinationSum2(candidates, target))
	}
}

//------------------------Leetcode Problem 40------------------------
/*
 * https://leetcode.cn/problems/combination-sum-ii/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.4MB 在所有Go提交中击败了87.68%的用户
**/
