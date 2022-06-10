package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	scores := make([]float64, n)
	var max, min, ave float64 = 0, 100, 0
	for i := 0; i < n; i++ {
		fmt.Scan(&scores[i])
		max = math.Max(max, scores[i])
		min = math.Min(min, scores[i])
		ave += scores[i]
	}
	fmt.Printf("%.2f %.2f %.2f", max, min, ave/float64(n))
}

/*
 * 运行时间：10ms 超过0.00%用Go提交的代码
 * 占用内存：836KB 超过75.00%用Go提交的代码
**/
