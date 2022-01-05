package main

import (
	"bufio"
	"os"
)

func modifyString(s string) string {
	arr, n := []int32(s), len(s)
	for i := range arr {
		if arr[i] == '?' {
			for ch := 'a'; ch <= 'c'; ch++ {
				if (i > 0 && arr[i-1] == ch) || (i < n-1 && arr[i+1] == ch) {
					continue
				}
				arr[i] = ch
				break
			}
		}
	}
	return string(arr)
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		println(modifyString(input.Text()))
	}
}

/*
 * 执行用时：4ms 在所有Go提交中击败了44.83%的用户
 * 占用内存：2.2MB 在所有Go提交中击败了10.34%的用户
**/
