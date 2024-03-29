package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"sort"
	"strconv"
)

// https://leetcode.cn/problems/cut-off-trees-for-golf-event/
//------------------------Leetcode Problem 675------------------------
func cutOffTree(forest [][]int) int {
	dirs := []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	type pair struct{ dis, x, y int }
	var trees []pair
	for i, row := range forest {
		for j, x := range row {
			if x > 1 {
				trees = append(trees, pair{x, i, j})
			}
		}
	}
	sort.Slice(trees, func(i, j int) bool { return trees[i].dis < trees[j].dis })
	bfs := func(sx, sy, tx, ty int) int {
		m, n := len(forest), len(forest[0])
		vis := make([][]bool, m)
		for i := range vis {
			vis[i] = make([]bool, n)
		}
		vis[sx][sy] = true
		q := []pair{{0, sx, sy}}
		for len(q) > 0 {
			p := q[0]
			q = q[1:]
			if p.x == tx && p.y == ty {
				return p.dis
			}
			for _, d := range dirs {
				if x, y := p.x+d.x, p.y+d.y; x >= 0 && x < m && y >= 0 && y < n && !vis[x][y] && forest[x][y] > 0 {
					vis[x][y] = true
					q = append(q, pair{p.dis + 1, x, y})
				}
			}
		}
		return -1
	}
	var preX, preY, ans int
	for _, t := range trees {
		d := bfs(preX, preY, t.x, t.y)
		if d < 0 {
			return -1
		}
		ans += d
		preX, preY = t.x, t.y
	}
	return ans
}

//------------------------Leetcode Problem 675------------------------
/*
 * https://leetcode.cn/problems/cut-off-trees-for-golf-event/
 * 执行用时：156ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：7.1MB 在所有Go提交中击败了85.71%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		forest := make([][]int, n)
		for i := range forest {
			scanner.Scan()
			forest[i] = utils.StringToInts(scanner.Text())
		}
		Printf("Output: %v\n", cutOffTree(forest))
	}
}
