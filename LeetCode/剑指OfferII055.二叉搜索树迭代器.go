package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	Arr []int
}

func (it *BSTIterator) Inorder(node *TreeNode) {
	if node == nil {
		return
	}
	it.Inorder(node.Left)
	it.Arr = append(it.Arr, node.Val)
	it.Inorder(node.Right)
}

func Constructor(root *TreeNode) (it BSTIterator) {
	it.Inorder(root)
	return
}

func (it *BSTIterator) Next() int {
	val := it.Arr[0]
	it.Arr = it.Arr[1:]
	return val
}

func (it *BSTIterator) HasNext() bool {
	return len(it.Arr) > 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */

/*
 * 执行用时：20ms 在所有Go提交中击败了88.24%的用户
 * 占用内存：9.7MB 在所有Go提交中击败了35.83%的用户
**/
