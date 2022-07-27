package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/minimum-jumps-to-reach-home/
//------------------------Leetcode Problem 1654------------------------
func minimumJumps(forbidden []int, a int, b int, x int) int {
	type queue struct {
		pos  int
		back bool
	}
	path := make(map[int]bool)
	var m int
	for _, x := range forbidden {
		m = max(m, x)
		path[x] = true
	}
	m += a + b + x
	path[0] = true
	q := []queue{{0, false}}
	var ans int
	for len(q) != 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			t := q[i]
			if t.pos == x {
				return ans
			}
			if t.pos+a <= m && !path[t.pos+a] {
				q = append(q, queue{t.pos + a, false})
				path[t.pos+a] = true
			}
			if !t.back && t.pos-b > 0 && !path[t.pos-b] {
				q = append(q, queue{t.pos - b, true})
			}
		}
		ans++
		q = q[size:]
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//------------------------Leetcode Problem 1654------------------------
/*
 * https://leetcode.cn/problems/minimum-jumps-to-reach-home/
 * 执行用时：8ms 在所有Go提交中击败了84.09%的用户
 * 占用内存：6.6MB 在所有Go提交中击败了61.36%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		forbidden := utils.StringToInts(scanner.Text())
		scanner.Scan()
		var a, b, x int
		Sscanf(scanner.Text(), "%d %d %d", &a, &b, &x)
		Printf("Output: %v\n", minimumJumps(forbidden, a, b, x))
	}
}
