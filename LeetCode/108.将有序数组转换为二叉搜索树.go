package main

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	var damn func([]int, int, int) *TreeNode
	damn = func(nums []int, left int, right int) *TreeNode {
		if right < left {
			return nil
		}
		mid := left + ((right - left) >> 1)
		root := &TreeNode{nums[mid], damn(nums, left, mid-1), damn(nums, mid+1, right)}
		return root
	}
	return damn(nums, 0, len(nums)-1)
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：3.3MB 在所有Go提交中击败了68.33%的用户
**/
