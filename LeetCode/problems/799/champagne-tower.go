package main

import (
	"bufio"
	. "fmt"
	"os"
)

// https://leetcode.cn/problems/champagne-tower/
//------------------------Leetcode Problem 799------------------------
func champagneTower(poured int, query_row int, query_glass int) float64 {
	row := []float64{float64(poured)}
	for i := 1; i <= query_row; i++ {
		nextRow := make([]float64, i+1)
		for j, volume := range row {
			if volume > 1 {
				nextRow[j] += (volume - 1) / 2
				nextRow[j+1] += (volume - 1) / 2
			}
		}
		row = nextRow
	}
	if row[query_glass] < 1 {
		return row[query_glass]
	}
	return 1
}

//------------------------Leetcode Problem 799------------------------
/*
 * https://leetcode.cn/problems/champagne-tower/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：4.7MB 在所有Go提交中击败了71.43%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var poured, query_row, query_glass int
		Sscanf(scanner.Text(), "%d %d %d", &poured, &query_row, &query_glass)
		Printf("Output: %v\n", champagneTower(poured, query_row, query_glass))
	}
}
