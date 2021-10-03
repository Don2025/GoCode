/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) int {
    if root == nil {
        return 0
    }
    return rootSum(root, targetSum)+pathSum(root.Left, targetSum)+pathSum(root.Right, targetSum)
}

func rootSum(root *TreeNode, targetSum int) int {
    if root == nil {
        return 0
    }
    var ans int
    if root.Val == targetSum {
        ans++
    }
    ans += rootSum(root.Left, targetSum-root.Val)+rootSum(root.Right, targetSum-root.Val)
    return ans
}
