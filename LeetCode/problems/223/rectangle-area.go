package main

import (
	. "fmt"
)

// https://leetcode.cn/problems/rectangle-area/
//------------------------Leetcode Problem 223------------------------
func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	area1, area2 := (ax2-ax1)*(ay2-ay1), (bx2-bx1)*(by2-by1)
	if ax2 <= bx1 || ay2 <= by1 || ax1 >= bx2 || ay1 >= by2 { //两矩阵无重叠部分
		return area1 + area2
	}
	topX, topY, bottomX, bottomY := min(ax2, bx2), min(ay2, by2), max(ax1, bx1), max(ay1, by1)
	return area1 + area2 - (topX-bottomX)*(topY-bottomY)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//------------------------Leetcode Problem 223------------------------
/*
 * https://leetcode.cn/problems/rectangle-area/
 * 执行用时：12ms 在所有Go提交中击败了85.71%的用户
 * 占用内存：6.2MB 在所有Go提交中击败了100.00%的用户
**/

func main() {
	for {
		var ax1, ay1, ax2, ay2, bx1, by1, bx2, by2 int
		Scanf("%d %d %d %d %d %d %d %d", &ax1, &ay1, &ax2, &ay2, &bx1, &by1, &bx2, &by2)
		Printf("Output: %v\n", computeArea(ax1, ay1, ax2, ay2, bx1, by1, bx2, by2))
	}
}
