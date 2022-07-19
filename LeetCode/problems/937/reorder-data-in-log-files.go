package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
	"strings"
)

// https://leetcode.cn/problems/reorder-data-in-log-files/
//------------------------Leetcode Problem 937------------------------
func reorderLogFiles(logs []string) []string {
	sort.SliceStable(logs, func(i, j int) bool {
		l1 := strings.Split(logs[i], " ")
		letterI := l1[1][0] < '0' || l1[1][0] > '9'
		l2 := strings.Split(logs[j], " ")
		letterJ := l2[1][0] < '0' || l2[1][0] > '9'
		// 字母日志都排在数字日志之前
		if letterI != letterJ {
			return letterI
		}
		// 数字日志按照原序排列
		if !letterI && !letterJ {
			return i < j
		}
		sI := strings.Join(l1[1:], " ")
		sJ := strings.Join(l2[1:], " ")
		// 字母日志在内容不同时, 忽略标识符后按内容字母顺序排序; 在内容相同时按标识符排序
		return sI < sJ || (sI == sJ && l1[0] < l2[0])
	})
	return logs
}

//------------------------Leetcode Problem 937------------------------
/*
 * https://leetcode.cn/problems/reorder-data-in-log-files/
 * 执行用时：4ms 在所有Go提交中击败了83.33%的用户
 * 占用内存：6.4MB 在所有Go提交中击败了12.50%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		logs := strings.Fields(scanner.Text())
		ans := reorderLogFiles(logs)
		Printf("Output: %v\n", ans)
	}
}
