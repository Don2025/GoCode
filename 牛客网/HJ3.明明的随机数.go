package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		var hash [1001]bool
		n, _ := strconv.Atoi(input.Text())
		for i := 0; i < n; i++ {
			var x int
			input.Scan()
			x, _ = strconv.Atoi(input.Text())
			hash[x] = true
		}
		for i, j := range hash {
			if j == true {
				fmt.Printf("%d\n", i)
			}
		}
	}
}

/*
 * 运行时间：2ms 超过100.00%用Go提交的代码
 * 占用内存：904KB 超过72.62%用Go提交的代码
**/
