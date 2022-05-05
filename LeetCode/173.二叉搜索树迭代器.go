package main

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

type BSTIterator struct {
	arr []int
}

func Constructor(root *TreeNode) (it BSTIterator) {
	it.Inorder(root)
	return
}

func (it *BSTIterator) Inorder(node *TreeNode) {
	if node == nil {
		return
	}
	it.Inorder(node.Left)
	it.arr = append(it.arr, node.Val)
	it.Inorder(node.Right)
}

func (it *BSTIterator) Next() int {
	val := it.arr[0]
	it.arr = it.arr[1:]
	return val
}

func (it *BSTIterator) HasNext() bool {
	return len(it.arr) > 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */

/*
 * 执行用时：24ms 在所有Go提交中击败了68.71%的用户
 * 占用内存：9.7MB 在所有Go提交中击败了45.71%的用户
**/
