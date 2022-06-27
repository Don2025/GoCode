package _97

import (
	"github.com/Don2025/GoCode/structures"
	"strconv"
	"strings"
)

type TreeNode = structures.TreeNode

// https://leetcode.cn/problems/serialize-and-deserialize-binary-tree/
//------------------------Leetcode Problem 297------------------------

type Codec struct{}

func Constructor() (_ Codec) {
	return
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	sb := &strings.Builder{}
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			sb.WriteString("null,")
			return
		}
		sb.WriteString(strconv.Itoa(node.Val))
		sb.WriteByte(',')
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return sb.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	sp := strings.Split(data, ",")
	var build func() *TreeNode
	build = func() *TreeNode {
		if sp[0] == "null" {
			sp = sp[1:]
			return nil
		}
		val, _ := strconv.Atoi(sp[0])
		sp = sp[1:]
		return &TreeNode{val, build(), build()}
	}
	return build()
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */

//------------------------Leetcode Problem 297------------------------
/*
 * https://leetcode.com/problems/serialize-and-deserialize-binary-tree/
 * 执行用时：8ms 在所有Go提交中击败了85.61%的用户
 * 占用内存：6.3MB 在所有Go提交中击败了91.91%的用户
**/
