package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		text := input.Text()
		split := strings.Split(text, " ")
		arr := stringArrayToIntArray(split)
		var maxans int
		for _, x := range arr {
			if maxans < x {
				maxans = x
			}
		}
		fmt.Println(maxans)
	}
}

func stringArrayToIntArray(strings []string) []int {
	var arr []int
	for _, x := range strings {
		n, _ := strconv.Atoi(x)
		arr = append(arr, n)
	}
	return arr
}

/*
 * 运行时间：4ms 超过50.00%用Go提交的代码
 * 占用内存：984KB 超过0.00%用Go提交的代码
**/
