package main

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func maxValue(root *TreeNode, k int) int {
	var dp func(*TreeNode) []int
	dp = func(node *TreeNode) []int {
		ans := make([]int, k+1)
		if node == nil {
			return ans
		}
		l, r := dp(node.Left), dp(node.Right)
		for i := 0; i < k; i++ {
			for j := 0; j < k-i; j++ {
				ans[i+j+1] = max(ans[i+j+1], l[i]+r[j]+node.Val)
			}
		}
		for i := 0; i <= k; i++ {
			for j := 0; j <= k; j++ {
				ans[0] = max(ans[0], l[i]+r[j])
			}
		}
		return ans
	}
	ans := dp(root)
	largest := ans[0]
	for i := 1; i < len(ans); i++ {
		if largest < ans[i] {
			largest = ans[i]
		}
	}
	return largest
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
 * 执行用时：68ms 在所有Go提交中击败了50.00%的用户
 * 占用内存：9.7MB 在所有Go提交中击败了60.00%的用户
**/
