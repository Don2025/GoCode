package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func maximumDetonation(bombs [][]int) int {
	n := len(bombs)
	g := make([][]int, n)
	//根据能否引爆旁边的炸弹建立邻接表
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if pow2(bombs[j][0]-bombs[i][0])+pow2(bombs[j][1]-bombs[i][1]) <= pow2(bombs[i][2]) {
				g[i] = append(g[i], j)
			}
		}
	}
	var bfs func(int) int
	bfs = func(i int) int {
		queue := []int{i}
		vis := make([]bool, n)
		vis[i] = true
		cnt := 1
		for len(queue) > 0 {
			x := queue[0]
			queue = queue[1:]
			for _, y := range g[x] {
				if !vis[y] {
					vis[y] = true
					cnt++
					queue = append(queue, y)
				}
			}
		}
		return cnt
	}
	var ans int
	//广度优先遍历每个结点找出最大的引爆数量
	for i := 0; i < n; i++ {
		ans = max(ans, bfs(i))
	}
	return ans
}

func pow2(n int) int {
	return n * n
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		bombs := make([][]int, n)
		for i := 0; i < n; i++ {
			input.Scan()
			bombs[i] = stringArrayToIntArray(strings.Fields(input.Text()))
		}
		println(maximumDetonation(bombs))
	}
}

func stringArrayToIntArray(strings []string) []int {
	var arr []int
	for _, x := range strings {
		n, _ := strconv.Atoi(x)
		arr = append(arr, n)
	}
	return arr
}

/*
 * 执行用时：24ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：6.9MB 在所有Go提交中击败了100.00%的用户
**/
