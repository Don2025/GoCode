package main

import "fmt"

func main() {
	const n, m = 5, 6
	stu := make([][]float64, n)
	for i := 0; i < n; i++ {
		stu[i] = make([]float64, m)
		for j := 0; j < n; j++ {
			fmt.Scan(&stu[i][j])
			fmt.Printf("%.1f ", stu[i][j])
			stu[i][n] += stu[i][j]
		}
		fmt.Printf("%.1f\n", stu[i][n])
	}
}

/*
 * 运行时间：3ms 超过100.00%用Go提交的代码
 * 占用内存：984KB 超过0.00%用Go提交的代码
**/
