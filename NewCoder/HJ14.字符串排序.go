package main

import (
	"fmt"
	"sort"
)

func main() {
	var arr []string
	var n int
	fmt.Scanf("%d", &n)
	for i := 1; i <= n; i++ {
		var s string
		fmt.Scanf("%s", &s)
		arr = append(arr, s)
	}
	sort.Strings(arr)
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%s\n", arr[i])
	}
}

/*
 * 运行时间：97ms 超过0.20%用Go提交的代码
 * 占用内存：1180KB 超过9.32%用Go提交的代码
**/
