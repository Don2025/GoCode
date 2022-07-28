package main

import "github.com/Don2025/GoCode/structures"

type ListNode = structures.ListNode

// https://leetcode.cn/problems/c32eOV/
// ------------------------剑指 Offer II Problem 22------------------------
func detectCycle(head *ListNode) *ListNode {
	m := make(map[*ListNode]bool)
	for p := head; p != nil; p = p.Next {
		if _, ok := m[p]; ok {
			return p
		}
		m[p] = true
	}
	return nil
}

// ------------------------剑指 Offer II Problem 22------------------------
/*
 * https://leetcode.cn/problems/c32eOV/
 * 执行用时：8ms 在所有Go提交中击败了24.60%的用户
 * 占用内存：4.9MB 在所有Go提交中击败了9.90%的用户
**/
