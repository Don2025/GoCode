package main

import (
	"fmt"
	"time"
)

func main() {
	const n = 10000
	for i := 0; i < n; i++ {
		go goFun(i) //开启一个并发携程
	}
	time.Sleep(time.Second)
}

func goFun(i int) {
	fmt.Println("goroutine ", i, "...")
}
