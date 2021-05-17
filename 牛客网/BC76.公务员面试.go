package main

import (
	"fmt"
	"sort"
)

func main() {
	var a [7]int
	for i := 0; i < 7; i++ {
		_, err := fmt.Scan(&a[i])
		if err != nil {
			return
		}
	}
	sort.Ints(a[:])
	sum := 0
	for i := 1; i < 6; i++ {
		sum += a[i]
	}
	fmt.Printf("%.2f", float64(sum)/5)
}

/*
 * 运行时间：9ms 超过22.22%用Go提交的代码
 * 占用内存：1080KB 超过0.00%用Go提交的代码
**/
