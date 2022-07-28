package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/structures"
	"github.com/Don2025/GoCode/utils"
	"os"
)

type ListNode = structures.ListNode

// https://leetcode.cn/problems/liang-ge-lian-biao-de-di-yi-ge-gong-gong-jie-dian-lcof/
// ------------------------剑指 Offer I Problem 52------------------------
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	//这两个链表就像是两个人的人生轨迹，若是有缘总会相遇
	if headA == nil || headB == nil { //如果我们不曾相遇，人生没有交集，那么也就没有以后啦
		return nil
	}
	you, me := headA, headB //你和我相约在未来某天一定会相遇携手到终点
	for you != me {         //某年某月的某一天就是对的时间遇到对的一切
		if you != nil { //你需要独自度过漫长的人生旅程
			you = you.Next
		} else { //你会不会突然的出现
			you = headB //那一天那一刻那个场景，你出现在我生命
		}
		if me != nil { //我独自走完必须要走的人生道路
			me = me.Next
		} else { //才会开始了解你曾经走过的那些路，渐渐地爱上你喜欢的一切
			me = headA //我来到你的城市，走过你来时的路
		}
	} //无数时间线，无尽可能性，终于交织向你
	return you //未来多漫长再漫长还有期待，陪伴你一直到故事给说完
}

// ------------------------剑指 Offer I Problem 52------------------------
/*
 * https://leetcode.cn/problems/liang-ge-lian-biao-de-di-yi-ge-gong-gong-jie-dian-lcof/
 * 执行用时：40ms 在所有Go提交中击败了84.47%的用户
 * 占用内存：9.2MB 在所有Go提交中击败了40.69%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		headA := utils.StringToListNode(scanner.Text())
		scanner.Scan()
		headB := utils.StringToListNode(scanner.Text())
		Printf("Output: %v\n", getIntersectionNode(headA, headB))
	}
}
