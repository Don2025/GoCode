package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
	"strings"
)

// https://leetcode.cn/problems/exclusive-time-of-functions/
//------------------------Leetcode Problem 636------------------------
func exclusiveTime(n int, logs []string) []int {
	ans := make([]int, n)
	type Pair struct{ idx, timestamp int }
	var stack []Pair
	for _, log := range logs {
		arr := strings.Split(log, ":")
		idx, _ := strconv.Atoi(arr[0])
		timestamp, _ := strconv.Atoi(arr[2])
		if arr[1] == "start" {
			if len(stack) > 0 {
				ans[stack[len(stack)-1].idx] += timestamp - stack[len(stack)-1].timestamp
				stack[len(stack)-1].timestamp = timestamp
			}
			stack = append(stack, Pair{idx, timestamp})
		} else {
			ans[stack[len(stack)-1].idx] += timestamp - stack[len(stack)-1].timestamp + 1
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				stack[len(stack)-1].timestamp = timestamp + 1
			}
		}
	}
	return ans
}

//------------------------Leetcode Problem 636------------------------
/*
 * https://leetcode.cn/problems/exclusive-time-of-functions/
 * 执行用时：8ms 在所有Go提交中击败了69.57%的用户
 * 占用内存：6.2MB 在所有Go提交中击败了26.09%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		logs := strings.Fields(scanner.Text())
		Printf("Output: %v\n", exclusiveTime(n, logs))
	}
}
