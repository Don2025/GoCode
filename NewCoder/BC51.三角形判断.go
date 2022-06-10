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
	isStart := true
	for input.Scan() {
		if isStart {
			isStart = false
		} else {
			fmt.Printf("\n")
		}
		text := input.Text()
		split := strings.Split(text, " ")
		arr := StringArrayToIntArray(split)
		if arr[0]+arr[1] > arr[2] && arr[0]+arr[2] > arr[1] && arr[2]+arr[1] > arr[0] {
			if arr[0] == arr[1] && arr[1] == arr[2] {
				fmt.Print("Equilateral triangle!")
			} else if arr[0] != arr[1] && arr[1] != arr[2] && arr[2] != arr[0] {
				fmt.Print("Ordinary triangle!")
			} else {
				fmt.Print("Isosceles triangle!")
			}
		} else {
			fmt.Print("Not a triangle!")
		}

	}
}

func StringArrayToIntArray(strings []string) []int {
	var arr []int
	for _, x := range strings {
		n, _ := strconv.Atoi(x)
		arr = append(arr, n)
	}
	return arr
}

/*
 * 运行时间：10ms 超过0.00%用Go提交的代码
 * 占用内存：960KB 超过11.11%用Go提交的代码
**/
