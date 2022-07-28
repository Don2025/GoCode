package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/count-nodes-with-the-highest-score/
//------------------------Leetcode Problem 2049------------------------
func countHighestScoreNodes(parents []int) int {
	n := len(parents)
	children := make([][]int, n)
	for i := 1; i < len(children); i++ {
		p := parents[i]
		children[p] = append(children[p], i)
	}
	var ans, maxScore int
	var dfs func(int) int
	dfs = func(i int) int {
		score, size := 1, n-1
		for _, x := range children[i] {
			sz := dfs(x)
			score *= sz
			size -= sz
		}
		if i > 0 {
			score *= size
		}
		if score == maxScore {
			ans++
		} else if score > maxScore {
			maxScore = score
			ans = 1
		}
		return n - size
	}
	dfs(0)
	return ans
}

//------------------------Leetcode Problem 2049------------------------
/*
 * https://leetcode.cn/problems/count-nodes-with-the-highest-score/
 * 执行用时：136ms 在所有Go提交中击败了82.76%的用户
 * 占用内存：29.7MB 在所有Go提交中击败了17.24%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		parents := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", countHighestScoreNodes(parents))
	}
}
