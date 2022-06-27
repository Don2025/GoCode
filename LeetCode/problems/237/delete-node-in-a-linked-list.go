package main

import (
	"bufio"
	"fmt"
	"github.com/Don2025/GoCode/structures"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

type ListNode = structures.ListNode

// https://leetcode.cn/problems/delete-node-in-a-linked-list/
//------------------------Leetcode Problem 237------------------------
func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

//------------------------Leetcode Problem 237------------------------
/*
 * https://leetcode.cn/problems/delete-node-in-a-linked-list/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.8MB 在所有Go提交中击败了55.79%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		head := utils.StringToListNode(scanner.Text())
		utils.PrintLinkedList(head)
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())
		for p := head; p != nil; p = p.Next {
			if p.Val == n {
				node := p
				deleteNode(node)
			}
		}
		fmt.Printf("After deleteNode:\n")
		utils.PrintLinkedList(head)
	}
}
