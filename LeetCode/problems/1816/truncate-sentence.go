package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// https://leetcode.cn/problems/truncate-sentence/
//------------------------Leetcode Problem 1816------------------------
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

//------------------------Leetcode Problem 1816------------------------
/*
 * https://leetcode.cn/problems/truncate-sentence/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.2MB 在所有Go提交中击败了28.57%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()
		scanner.Scan()
		k, _ := strconv.Atoi(scanner.Text())
		println(truncateSentence(s, k))
	}
}
