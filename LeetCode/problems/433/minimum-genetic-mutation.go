package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
)

// https://leetcode.cn/problems/minimum-genetic-mutation/
//------------------------Leetcode Problem 433------------------------
func minMutation(start string, end string, bank []string) int {
	type queue struct {
		str  string
		step int
	}
	var only1Different func(string, string) bool
	only1Different = func(a string, b string) bool {
		var n int
		for i := range a {
			if a[i] != b[i] {
				n++
			}
			if n > 1 {
				break
			}
		}
		return n == 1
	}
	visited := make([]bool, len(bank))
	q := []queue{{start, 0}}
	for len(q) != 0 {
		t := q[0]
		q = q[1:]
		if t.str == end {
			return t.step
		}
		for i := 0; i < len(bank); i++ {
			if visited[i] || !only1Different(t.str, bank[i]) {
				continue
			}
			q = append(q, queue{bank[i], t.step + 1})
			visited[i] = true
		}
	}
	return -1
}

//------------------------Leetcode Problem 433------------------------
/*
 * https://leetcode.cn/problems/minimum-genetic-mutation/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了64.66%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		start := scanner.Text()
		scanner.Scan()
		end := scanner.Text()
		scanner.Scan()
		bank := strings.Fields(scanner.Text())
		Printf("Output: %v\n", minMutation(start, end, bank))
	}
}
