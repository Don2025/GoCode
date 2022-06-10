package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param head ListNode类
 * @return int整型一维数组
 */
func printListFromTailToHead(head *ListNode) []int {
	ans := []int{}
	var visit func(node *ListNode)
	visit = func(node *ListNode) {
		if node == nil {
			return
		}
		visit(node.Next)
		ans = append(ans, node.Val)
	}
	visit(head)
	return ans
}

/*
 * 运行时间：2ms 超过100.00%用Go提交的代码
 * 占用内存：820KB 超过56.60%用Go提交的代码
**/
