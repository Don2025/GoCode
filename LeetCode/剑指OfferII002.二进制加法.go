package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func addBinary(a string, b string) string {
	var ans string
	var carry int
	na, nb := len(a), len(b)
	n := max(na, nb)
	for i := 0; i < n; i++ {
		if i < na {
			carry += int(a[na-1-i] - '0')
		}
		if i < nb {
			carry += int(b[nb-1-i] - '0')
		}
		ans = strconv.Itoa(carry%2) + ans
		carry >>= 1
	}
	if carry > 0 {
		ans = "1" + ans
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		arr := strings.Split(input.Text(), "+")
		println(addBinary(arr[0], arr[1]))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.2MB 在所有Go提交中击败了100.00%的用户
**/
