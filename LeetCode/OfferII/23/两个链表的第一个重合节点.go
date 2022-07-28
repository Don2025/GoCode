package main

import "github.com/Don2025/GoCode/structures"

type ListNode = structures.ListNode

// https://leetcode.cn/problems/3u1WK4/
// ------------------------剑指 Offer II Problem 23------------------------
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pa, pb := headA, headB
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}
		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pa
}

// ------------------------剑指 Offer II Problem 23------------------------
/*
 * https://leetcode.cn/problems/3u1WK4/
 * 执行用时：28ms 在所有Go提交中击败了74.80%的用户
 * 占用内存：7MB 在所有Go提交中击败了87.70%的用户
**/
