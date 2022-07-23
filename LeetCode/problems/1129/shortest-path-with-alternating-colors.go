package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"math"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/shortest-path-with-alternating-colors/
//------------------------Leetcode Problem 1129------------------------
func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
	redG := make([][]int, n)
	for _, e := range redEdges {
		redG[e[0]] = append(redG[e[0]], e[1])
	}
	blueG := make([][]int, n)
	for _, e := range blueEdges {
		blueG[e[0]] = append(blueG[e[0]], e[1])
	}
	redV := make([]bool, n)
	blueV := make([]bool, n)
	type queue struct {
		x int
		t bool
	} // t==true means red, false means blue
	q := []queue{{0, true}, {0, false}}
	var d int
	ans := make([]int, n)
	for i := range ans {
		ans[i] = math.MaxInt32
	}
	for len(q) != 0 {
		i := len(q)
		for ; i > 0; i-- {
			t := q[0]
			q = q[1:]
			ans[t.x] = min(ans[t.x], d)
			if t.t { // find red edge
				for _, v := range redG[t.x] {
					if !blueV[v] {
						blueV[v] = true
						q = append(q, queue{v, !t.t})
					}
				}
			} else { // find blue edge
				for _, v := range blueG[t.x] {
					if !redV[v] {
						redV[v] = true
						q = append(q, queue{v, !t.t})
					}
				}
			}
		}
		d++
	}
	for i := range ans {
		if ans[i] == math.MaxInt32 {
			ans[i] = -1
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//------------------------Leetcode Problem 1129------------------------
/*
 * https://leetcode.cn/problems/shortest-path-with-alternating-colors/
 * 执行用时：8ms 在所有Go提交中击败了81.82%的用户
 * 占用内存：5.1MB 在所有Go提交中击败了77.27%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		redN, _ := strconv.Atoi(scanner.Text())
		red := make([][]int, redN)
		for i := 0; i < redN; i++ {
			scanner.Scan()
			red[i] = utils.StringToInts(scanner.Text())
		}
		scanner.Scan()
		blueN, _ := strconv.Atoi(scanner.Text())
		blue := make([][]int, blueN)
		for i := 0; i < blueN; i++ {
			scanner.Scan()
			blue[i] = utils.StringToInts(scanner.Text())
		}
		Printf("Output: %v\n", shortestAlternatingPaths(n, red, blue))
	}
}
