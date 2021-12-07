package main

import (
	"fmt"
	"sort"
)

func main() {
	for {
		a := make([]int, 7)
		for i := 0; i < 7; i++ {
			_, err := fmt.Scan(&a[i])
			if err != nil {
				return
			}
		}
		sort.Ints(a)
		var sum int
		for i := 1; i < 6; i++ {
			sum += a[i]
		}
		fmt.Printf("%.2f\n", float64(sum)/5)
	}
}

/*
 * 运行时间：3ms 超过50.00%用Go提交的代码
 * 占用内存：1088KB 超过0.00%用Go提交的代码
**/
