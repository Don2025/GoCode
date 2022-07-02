package main

import (
	"bufio"
	. "fmt"
	"math"
	"os"
	"strings"
)

// https://leetcode.cn/problems/minimum-index-sum-of-two-lists/
//------------------------Leetcode Problem 599------------------------
func findRestaurant(list1 []string, list2 []string) []string {
	m := make(map[string]int, len(list1))
	for i, x := range list1 {
		m[x] = i
	}
	sum := math.MaxInt32
	var ans []string
	for i, x := range list2 {
		if j, ok := m[x]; ok {
			if i+j < sum {
				sum = i + j
				ans = []string{x}
			} else if i+j == sum {
				ans = append(ans, x)
			}
		}
	}
	return ans
}

//------------------------Leetcode Problem 599------------------------
/*
 * https://leetcode.cn/problems/minimum-index-sum-of-two-lists/
 * 执行用时：24ms 在所有Go提交中击败了86.96%的用户
 * 占用内存：6.8MB 在所有Go提交中击败了100.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		list1 := strings.Fields(scanner.Text())
		scanner.Scan()
		list2 := strings.Fields(scanner.Text())
		Printf("Output: %v\n", findRestaurant(list1, list2))
	}
}
