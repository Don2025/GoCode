package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/escape-the-ghosts/
//------------------------Leetcode Problem 789------------------------
func escapeGhosts(ghosts [][]int, target []int) bool {
	dis := abs(target[0]) + abs(target[1])
	for _, g := range ghosts {
		if abs(g[0]-target[0])+abs(g[1]-target[1]) <= dis {
			return false
		}
	}
	return true
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

//------------------------Leetcode Problem 789------------------------
/*
 * https://leetcode.cn/problems/escape-the-ghosts/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：0MB 在所有Go提交中击败了100.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		ghosts := make([][]int, n)
		for i := 0; i < n; i++ {
			scanner.Scan()
			ghosts[i] = utils.StringToInts(scanner.Text())
		}
		scanner.Scan()
		target := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", escapeGhosts(ghosts, target))
	}
}
