package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
	"strings"
)

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	n1, n2 := len(num1), len(num2)
	num := make([]int, n1+n2)
	for i := n1 - 1; i >= 0; i-- {
		x := int(num1[i]) - '0'
		for j := n2 - 1; j >= 0; j-- {
			y := int(num2[j]) - '0'
			num[i+j+1] += x * y
		}
	}
	for i := n1 + n2 - 1; i > 0; i-- {
		num[i-1] += num[i] / 10
		num[i] %= 10
	}
	var ans string
	var i int
	if num[0] == 0 {
		i = 1
	}
	for ; i < n1+n2; i++ {
		ans += strconv.Itoa(num[i])
	}
	return ans
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		arr := strings.Fields(scanner.Text())
		Println(multiply(arr[0], arr[1]))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.9MB 在所有Go提交中击败了78.54%的用户
**/
