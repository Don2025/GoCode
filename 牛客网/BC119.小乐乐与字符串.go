package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	var cCnt, hCnt, nCnt int
	for _, x := range s {
		switch x {
		case 'C':
			cCnt++
		case 'H':
			hCnt += cCnt
		case 'N':
			nCnt += hCnt
		}
	}
	fmt.Print(nCnt)
}

/*
 * 运行时间：7ms 超过0.00%用Go提交的代码
 * 占用内存：880KB 超过100.00%用Go提交的代码
**/
