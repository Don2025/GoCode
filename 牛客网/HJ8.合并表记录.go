package main

import (
	"fmt"
	"sort"
)

func main() {
	mp := make(map[int]int)
	var n, x, y int
	var a []int
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d %d", &x, &y)
		if mp[x] == 0 {
			a = append(a, x)
		}
		mp[x] += y
	}
	sort.Ints(a)
	for _, x := range a {
		if mp[x] != 0 {
			fmt.Printf("%d %d\n", x, mp[x])
		}
	}
}

/*
 * 运行时间：5ms 超过54.40%用Go提交的代码
 * 占用内存：912KB 超过72.70%用Go提交的代码
**/
