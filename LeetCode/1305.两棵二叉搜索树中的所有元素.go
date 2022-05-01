package main

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func inorder(root *TreeNode) (ans []int) {
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		ans = append(ans, root.Val)
		dfs(root.Right)
	}
	dfs(root)
	return
}

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	nums1, nums2 := inorder(root1), inorder(root2)
	p1, n1 := 0, len(nums1)
	p2, n2 := 0, len(nums2)
	merged := make([]int, 0, n1+n2)
	for {
		if p1 == n1 {
			return append(merged, nums2[p2:]...)
		}
		if p2 == n2 {
			return append(merged, nums1[p1:]...)
		}
		if nums1[p1] < nums2[p2] {
			merged = append(merged, nums1[p1])
			p1++
		} else {
			merged = append(merged, nums2[p2])
			p2++
		}
	}
}

/*
 * 执行用时：80ms 在所有Go提交中击败了86.21%的用户
 * 占用内存：8.4MB 在所有Go提交中击败了37.93%的用户
**/
