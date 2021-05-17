package main

import "fmt"

func main() {
	var s string
	_, err := fmt.Scan(&s)
	if err != nil {
		return
	}
	var na, nb int
	for _, it := range s {
		if it == 'A' {
			na++
		} else if it == 'B' {
			nb++
		}
	}
	if na > nb {
		fmt.Println("A")
	} else if na < nb {
		fmt.Println("B")
	} else {
		fmt.Println("E")
	}
}

/*
 * 运行时间：4ms 超过44.44%用Go提交的代码
 * 占用内存：880KB 超过55.56%用Go提交的代码
**/
