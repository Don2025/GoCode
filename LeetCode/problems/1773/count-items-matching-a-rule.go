package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
	"strings"
)

// https://leetcode.cn/problems/count-items-matching-a-rule/
//------------------------Leetcode Problem 1773------------------------
func countMatches(items [][]string, ruleKey string, ruleValue string) (cnt int) {
	rule := map[string]int{"type": 0, "color": 1, "name": 2}
	for _, item := range items {
		if item[rule[ruleKey]] == ruleValue {
			cnt++
		}
	}
	return
}

//------------------------Leetcode Problem 1773------------------------
/*
 * https://leetcode.cn/problems/count-items-matching-a-rule/
 * 执行用时：28ms 在所有Go提交中击败了86.36%的用户
 * 占用内存：6.8MB 在所有Go提交中击败了40.91%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		items := make([][]string, n)
		for i := range items {
			scanner.Scan()
			items[i] = strings.Fields(scanner.Text())
		}
		scanner.Scan()
		ruleKey := scanner.Text()
		scanner.Scan()
		ruleValue := scanner.Text()
		Printf("Output: %d\n", countMatches(items, ruleKey, ruleValue))
	}
}
