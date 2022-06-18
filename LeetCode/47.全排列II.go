package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"sort"
)

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	var ans [][]int
	var arr []int
	visited := make([]bool, n)
	var backtrack func(int)
	backtrack = func(index int) {
		if index == n {
			ans = append(ans, append([]int(nil), arr...))
			return
		}
		for i := range nums {
			if visited[i] || i > 0 && !visited[i-1] && nums[i] == nums[i-1] {
				continue
			}
			arr = append(arr, nums[i])
			visited[i] = true
			backtrack(index + 1)
			visited[i] = false
			arr = arr[:len(arr)-1]
		}
	}
	backtrack(0)
	return ans
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		ans := permuteUnique(nums)
		for _, x := range ans {
			Printf("%v ", x)
		}
		Println()
	}
}

/*
 * 执行用时：4ms 在所有Go提交中击败了62.35%的用户
 * 占用内存：3.6MB 在所有Go提交中击败了88.70%的用户
**/
