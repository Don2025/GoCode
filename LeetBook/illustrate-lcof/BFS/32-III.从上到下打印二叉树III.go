package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var ans [][]int
	if root == nil {
		return ans
	}
	q := []*TreeNode{0: root}
	var flag bool //false从左至右，true从右至左
	for len(q) != 0 {
		l := list.New()
		var a []int
		n := len(q)
		for i := 0; i < n; i++ {
			node := q[0]
			q = q[1:]
			if flag {
				l.PushFront(node.Val)
			} else {
				a = append(a, node.Val)
			}
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		if flag {
			for l.Len() != 0 {
				a = append(a, l.Front().Value.(int))
				l.Remove(l.Front())
			}
		}
		flag = !flag
		ans = append(ans, a)
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		root := createTree(strings.Fields(input.Text()), 0)
		ans := levelOrder(root)
		for _, v := range ans {
			for _, x := range v {
				fmt.Printf("%d ", x)
			}
			println()
		}
	}
}

func createTree(arr []string, index int) *TreeNode {
	p := new(TreeNode)
	if arr[index] == "null" {
		return nil
	}
	p.Val, _ = strconv.Atoi(arr[index])
	if index < len(arr) && (2*index+1) < len(arr) {
		p.Left = createTree(arr, 2*index+1)
	}
	if index < len(arr) && (2*index+2) < len(arr) {
		p.Right = createTree(arr, 2*index+2)
	}
	return p
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.9MB 在所有Go提交中击败了25.91%的用户
**/
