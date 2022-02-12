package main

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

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	var sum int
	for i := row1; i <= row2; i++ {
		sum += this.sums[i][col2+1] - this.sums[i][col1]
	}
	return sum
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */

/*
 * 执行用时：556ms 在所有Go提交中击败了84.23%的用户
 * 占用内存：19.3MB 在所有Go提交中击败了9.19%的用户
**/
