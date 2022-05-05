package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func checkIfExist(arr []int) bool {
	m := make(map[int]int, len(arr))
	for k, v := range arr {
		m[v] = k
	}
	for i := range arr {
		arr[i] <<= 1
	}
	for i, x := range arr {
		if j, ok := m[x]; i != j && ok {
			return true
		}
	}
	return false
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		arr := stringArrayToIntArray(strings.Fields(input.Text()))
		println(checkIfExist(arr))
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
 * 执行用时：4ms 在所有Go提交中击败了80s.99%的用户
 * 占用内存：3.6MB 在所有Go提交中击败了23.14%的用户
**/
