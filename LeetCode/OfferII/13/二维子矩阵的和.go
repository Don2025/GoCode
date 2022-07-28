package main

// https://leetcode.cn/problems/O4NDxx/
// ------------------------剑指 Offer II Problem 13------------------------

type NumMatrix struct {
	sums [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	sums := make([][]int, len(matrix))
	for i, row := range matrix {
		sums[i] = make([]int, len(row)+1)
		for j, x := range row {
			sums[i][j+1] = sums[i][j] + x
		}
	}
	return NumMatrix{sums}
}

func (nm *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	var sum int
	for i := row1; i <= row2; i++ {
		sum += nm.sums[i][col2+1] - nm.sums[i][col1]
	}
	return sum
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
// ------------------------剑指 Offer II Problem 13------------------------
/*
 * https://leetcode.cn/problems/O4NDxx/
 * 执行用时：252ms 在所有Go提交中击败了12.03%的用户
 * 占用内存：20.1MB 在所有Go提交中击败了5.16%的用户
**/
