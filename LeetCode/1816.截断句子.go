package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func truncateSentence(s string, k int) string {
	sb := strings.Builder{}
	arr := strings.Fields(s)
	for i := 0; i < k; i++ {
		if i > 0 {
			sb.WriteByte(' ')
		}
		sb.WriteString(arr[i])
	}
	return sb.String()
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		s := input.Text()
		input.Scan()
		k, _ := strconv.Atoi(input.Text())
		println(truncateSentence(s, k))
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.2MB 在所有Go提交中击败了28.57%的用户
**/
