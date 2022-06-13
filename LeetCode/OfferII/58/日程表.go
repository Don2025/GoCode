package main

import (
	"fmt"
)

// https://leetcode.cn/problems/fi9suh/
//------------------------剑指 Offer II Problem 58------------------------

type AVLNode struct {
	Left, Right *AVLNode
	Date        interface{}
	Height      int
}

func NewAVLNode(data interface{}) *AVLNode {
	return &AVLNode{nil, nil, data, 1}
}

func (avlNode *AVLNode) PrevNode(data interface{}) *AVLNode {
	if avlNode == nil {
		return nil
	}
	var found *AVLNode
	switch Compare(data, avlNode.Date) {
	case -1:
		return avlNode.Left.PrevNode(data)
	case 1:
		found = avlNode.Right.PrevNode(data)
		if found != nil {
			return found
		}
	case 0:
		return avlNode
	}
	return avlNode
}

func (avlNode *AVLNode) NextNode(data interface{}) *AVLNode {
	if avlNode == nil {
		return nil
	}
	var found *AVLNode
	switch Compare(data, avlNode.Date) {
	case -1:
		found = avlNode.Left.NextNode(data)
		if found != nil {
			return found
		}
	case 1:
		return avlNode.Right.NextNode(data)
	case 0:
		return avlNode
	}
	return avlNode
}

func Compare(left, right interface{}) int {
	a := left.([]int)[0]
	b := right.([]int)[0]
	if a > b {
		return 1
	} else if a < b {
		return -1
	} else {
		return 0
	}
}

// GetData gets node data
func (avlNode *AVLNode) GetData() interface{} {
	return avlNode.Date
}

// SetData sets node data
func (avlNode *AVLNode) SetData(data interface{}) {
	if avlNode == nil {
		return
	}
	avlNode.Date = data
}

func (avlNode *AVLNode) GetLeft() *AVLNode {
	if avlNode == nil {
		return nil
	}
	return avlNode.Left
}

func (avlNode *AVLNode) GetRight() *AVLNode {
	if avlNode == nil {
		return nil
	}
	return avlNode.Right
}

func (avlNode *AVLNode) GetHeight() int {
	if avlNode == nil {
		return 0
	}
	return avlNode.Height
}

// RightRotate LL type, clockwise rotation
func (avlNode *AVLNode) RightRotate() *AVLNode {
	headNode := avlNode.Left
	avlNode.Left = headNode.Right
	headNode.Right = avlNode
	avlNode.Height = Max(avlNode.Left.GetHeight(), avlNode.Right.GetHeight()) + 1
	headNode.Height = Max(headNode.Left.GetHeight(), headNode.Right.GetHeight()) + 1
	return headNode
}

// LeftRotate RR type, counterclockwise rotation
func (avlNode *AVLNode) LeftRotate() *AVLNode {
	headNode := avlNode.Right
	avlNode.Right = headNode.Left
	headNode.Left = avlNode
	avlNode.Height = Max(avlNode.Left.GetHeight(), avlNode.Right.GetHeight()) + 1
	headNode.Height = Max(headNode.Left.GetHeight(), headNode.Right.GetHeight()) + 1

	return headNode
}

// LeftThenRightRotate LR type, first rotate anticlockwise and then clockwise
func (avlNode *AVLNode) LeftThenRightRotate() *AVLNode {
	avlNode.Left = avlNode.Left.LeftRotate()
	return avlNode.RightRotate()
}

// RightThenLeftRotate RL type, first clockwise, then counterclockwise
func (avlNode *AVLNode) RightThenLeftRotate() *AVLNode {
	avlNode.Right = avlNode.Right.RightRotate()
	return avlNode.LeftRotate()
}

// adjust can adjust balance AVL tree
func (avlNode *AVLNode) adjust() *AVLNode {
	if avlNode.Right.GetHeight()-avlNode.Left.GetHeight() == 2 {
		if avlNode.Right.Right.GetHeight() > avlNode.Right.Left.GetHeight() {
			avlNode = avlNode.LeftRotate()
		} else {
			avlNode = avlNode.RightThenLeftRotate()
		}
	} else if avlNode.Right.GetHeight()-avlNode.Left.GetHeight() == -2 {
		if avlNode.Left.Left.GetHeight() > avlNode.Left.Right.GetHeight() {
			avlNode = avlNode.RightRotate()
		} else {
			avlNode = avlNode.LeftThenRightRotate()
		}
	}
	return avlNode
}

func (avlNode *AVLNode) Insert(data interface{}) *AVLNode {
	if avlNode == nil {
		return NewAVLNode(data)
	}
	switch Compare(data, avlNode.Date) {
	case -1:
		avlNode.Left = avlNode.Left.Insert(data)
		avlNode = avlNode.adjust()
	case 1:
		avlNode.Right = avlNode.Right.Insert(data)
		avlNode = avlNode.adjust()
	case 0:
		fmt.Print("The data already exists!")
	}
	avlNode.Height = Max(avlNode.Left.GetHeight(), avlNode.Right.GetHeight()) + 1
	return avlNode
}

func (avlNode *AVLNode) Find(data interface{}) *AVLNode {
	var found *AVLNode
	switch Compare(data, avlNode.Date) {
	case -1:
		found = avlNode.Left.Find(data)
	case 1:
		found = avlNode.Right.Find(data)
	case 0:
		return avlNode
	}
	return found
}

func (avlNode *AVLNode) FindMin() *AVLNode {
	var found *AVLNode
	if avlNode.Left != nil {
		found = avlNode.Left.FindMin()
	} else {
		found = avlNode
	}
	return found
}

func (avlNode *AVLNode) FindMax() *AVLNode {
	var found *AVLNode
	if avlNode.Right != nil {
		found = avlNode.Right.FindMax()
	} else {
		found = avlNode
	}
	return found
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type MyCalendar struct {
	Root *AVLNode
}

func Constructor() MyCalendar {
	var root *AVLNode
	return MyCalendar{
		Root: root,
	}
}

func (mc *MyCalendar) Book(start int, end int) bool {
	prevNode := mc.Root.PrevNode([]int{start, end})
	nextNode := mc.Root.NextNode([]int{start, end})
	if mc.Root == nil ||
		(prevNode == nil && end <= nextNode.Date.([]int)[0]) ||
		(nextNode == nil && prevNode.Date.([]int)[1] <= start) ||
		(prevNode != nil && nextNode != nil && end <= nextNode.Date.([]int)[0] && prevNode.Date.([]int)[1] <= start) {
		mc.Root = mc.Root.Insert([]int{start, end})
		return true
	}
	return false
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */

//------------------------剑指 Offer II Problem 58------------------------
/*
 * https://leetcode.cn/problems/fi9suh/
 * 执行用时：64ms 在所有Go提交中击败了82.71%的用户
 * 占用内存：7.1MB 在所有Go提交中击败了42.86%的用户
**/
