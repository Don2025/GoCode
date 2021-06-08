package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func lastRemaining(n, m int) int {
	var ans int
	for i := 2; i <= n; i++ {
		ans = (ans + m) % i
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		strArr := strings.Fields(input.Text())
		n, _ := strconv.Atoi(strArr[0])
		m, _ := strconv.Atoi(strArr[1])
		println(lastRemaining(n, m))
	}
}

/*
 * 执行用时：8ms 在所有Go提交中击败了99.53%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了87.88%的用户
**/
