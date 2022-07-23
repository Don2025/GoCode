package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/flower-planting-with-no-adjacent/
//------------------------Leetcode Problem 1042------------------------
func gardenNoAdj(n int, paths [][]int) []int {
	ans := make([]int, n)
	g := make([][]int, n)
	for _, x := range paths {
		g[x[0]-1] = append(g[x[0]-1], x[1]-1)
		g[x[1]-1] = append(g[x[1]-1], x[0]-1)
	}
	for i := 0; i < n; i++ {
		visited := make(map[int]bool)
		for _, x := range g[i] {
			if ans[x] != 0 {
				visited[ans[x]] = true
			}
		}
		for j := 1; j <= 4; j++ {
			if _, ok := visited[j]; !ok {
				ans[i] = j
				break
			}
		}
	}
	return ans
}

/*
 * https://leetcode.cn/problems/flower-planting-with-no-adjacent/
 * 执行用时：56ms 在所有Go提交中击败了60.00%的用户
 * 占用内存：8.4MB 在所有Go提交中击败了48.57%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		paths := make([][]int, n)
		for i := range paths {
			scanner.Scan()
			paths[i] = utils.StringToInts(scanner.Text())
		}
		Printf("Output: %v\n", gardenNoAdj(n, paths))
	}
}
