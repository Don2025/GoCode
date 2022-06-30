package main

import (
	"github.com/Don2025/GoCode/LeetCode"
	"github.com/Don2025/GoCode/structures"
	"math"
	"strconv"
	"strings"
)

type TreeNode = structures.TreeNode

// https://leetcode.cn/problems/serialize-and-deserialize-bst/
//------------------------Leetcode Problem 449------------------------
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct{}

func Constructor() (_ Codec) { return }

// Serializes a tree to a single string.
func (Codec) serialize(root *main.TreeNode) string {
	var arr []string
	var preOrder func(*main.TreeNode)
	preOrder = func(node *main.TreeNode) {
		if node == nil {
			return
		}
		arr = append(arr, strconv.Itoa(node.Val))
		preOrder(node.Left)
		preOrder(node.Right)
	}
	preOrder(root)
	return strings.Join(arr, " ")
}

// Deserializes your encoded data to tree.
func (Codec) deserialize(data string) *main.TreeNode {
	if data == "" {
		return nil
	}
	arr := strings.Split(data, " ")
	var dfs func(int, int) *main.TreeNode
	dfs = func(left, right int) *main.TreeNode {
		if len(arr) == 0 {
			return nil
		}
		val, _ := strconv.Atoi(arr[0])
		if val < left || val > right {
			return nil
		}
		arr = arr[1:]
		return &main.TreeNode{val, dfs(left, val), dfs(val, right)}
	}
	return dfs(math.MinInt32, math.MaxInt32)
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */
//------------------------Leetcode Problem 449------------------------
/*
 * https://leetcode.cn/problems/serialize-and-deserialize-bst/
 * 执行用时：8ms 在所有Go提交中击败了89.47%的用户
 * 占用内存：6.3MB 在所有Go提交中击败了81.05%的用户
**/
