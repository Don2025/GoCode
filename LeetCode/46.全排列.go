package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	}
	var ans [][]int
	var path []int
	visited := make([]bool, len(nums))
	var backtrack func([]int, []int, []bool)
	backtrack = func(nums, path []int, visited []bool) {
		if len(nums) == len(path) {
			t := make([]int, len(nums))
			copy(t, path)
			ans = append(ans, t)
			return
		}
		for i := range nums {
			if !visited[i] {
				visited[i] = true
				path = append(path, nums[i])
				backtrack(nums, path, visited)
				path = path[:len(path)-1]
				visited[i] = false
			}
		}
	}
	backtrack(nums, path, visited)
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		nums := stringArrayToIntArray(strings.Fields(input.Text()))
		ans := permute(nums)
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
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.7MB 在所有Go提交中击败了48.74%的用户
**/
