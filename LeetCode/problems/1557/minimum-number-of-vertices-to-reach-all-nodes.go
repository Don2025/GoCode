package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/minimum-number-of-vertices-to-reach-all-nodes/
//------------------------Leetcode Problem 1557------------------------
func findSmallestSetOfVertices(n int, edges [][]int) []int {
	var ans []int
	endSet := make(map[int]bool)
	for _, e := range edges {
		endSet[e[1]] = true
	}
	for i := 0; i < n; i++ {
		if !endSet[i] {
			ans = append(ans, i)
		}
	}
	return ans
}

//------------------------Leetcode Problem 1557------------------------
/*
 * https://leetcode.cn/problems/minimum-number-of-vertices-to-reach-all-nodes/
 * 执行用时：176ms 在所有Go提交中击败了37.33%的用户
 * 占用内存：18.4MB 在所有Go提交中击败了21.33%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		m, _ := strconv.Atoi(scanner.Text())
		edges := make([][]int, m)
		for i := 0; i < m; i++ {
			scanner.Scan()
			edges[i] = utils.StringToInts(scanner.Text())
		}
		Printf("Output: %v\n", findSmallestSetOfVertices(n, edges))
	}
}
