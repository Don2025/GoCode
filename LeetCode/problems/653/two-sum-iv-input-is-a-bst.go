package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/structures"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

type TreeNode = structures.TreeNode

// https://leetcode.cn/problems/two-sum-iv-input-is-a-bst/
//------------------------Leetcode Problem 653------------------------
func findTarget(root *TreeNode, k int) bool {
	set := make(map[int]bool)
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			continue
		}
		if set[k-node.Val] {
			return true
		}
		set[node.Val] = true
		queue = append(queue, node.Left)
		queue = append(queue, node.Right)
	}
	return false
}

//------------------------Leetcode Problem 653------------------------
/*
 * https://leetcode.cn/problems/two-sum-iv-input-is-a-bst/
 * 执行用时：16ms 在所有Go提交中击败了98.74%的用户
 * 占用内存：7.4MB 在所有Go提交中击败了60.92%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		root := utils.StringToTreeNode(scanner.Text())
		scanner.Scan()
		k, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", findTarget(root, k))
	}
}
