package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/structures"
	"os"
	"strconv"
)

type TreeNode = structures.TreeNode

// https://leetcode.cn/problems/unique-binary-search-trees-ii/
//------------------------Leetcode Problem 95------------------------
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	var helper func(int, int) []*TreeNode
	helper = func(start, end int) []*TreeNode {
		if start > end {
			return []*TreeNode{nil}
		}
		var ans []*TreeNode
		for i := start; i <= end; i++ {
			leftTrees := helper(start, i-1)
			rightTrees := helper(i+1, end)
			for _, left := range leftTrees {
				for _, right := range rightTrees {
					cur := &TreeNode{i, left, right}
					ans = append(ans, cur)
				}
			}
		}
		return ans
	}
	return helper(1, n)
}

//------------------------Leetcode Problem 95------------------------
/*
 * https://leetcode.cn/problems/unique-binary-search-trees-ii/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：4.2MB 在所有Go提交中击败了55.12%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", generateTrees(n))
	}
}
