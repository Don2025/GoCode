package main

import "fmt"

func main() {
	var user, passwd string
	for {
		_, err := fmt.Scan(&user, &passwd)
		if err != nil {
			return
		}
		if user == "admin" && passwd == "admin" {
			fmt.Println("Login Success!")
		} else {
			fmt.Println("Login Fail!")
		}
	}
}

/*
 * 运行时间：5ms 超过25.00%用Go提交的代码
 * 占用内存：1008KB 超过25.00%用Go提交的代码
**/
