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

// https://leetcode.cn/problems/construct-string-from-binary-tree/
//------------------------Leetcode Problem 606------------------------
func tree2str(root *TreeNode) string {
	var ans string
	if root == nil {
		return ans
	}
	ans = strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		return ans
	}
	if root.Left != nil {
		ans += "(" + tree2str(root.Left) + ")"
	} else {
		ans += "()"
	}
	if root.Right != nil {
		ans += "(" + tree2str(root.Right) + ")"
	}
	return ans
}

//------------------------Leetcode Problem 606------------------------
/*
 * https://leetcode.cn/problems/construct-string-from-binary-tree/
 * 执行用时：8ms 在所有Go提交中击败了87.40%的用户
 * 占用内存：7.2MB 在所有Go提交中击败了81.10%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		root := utils.StringToTreeNode(scanner.Text())
		Printf("Output: %v\n", tree2str(root))
	}
}
