package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

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

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		input.Scan()
		redN, _ := strconv.Atoi(input.Text())
		red := make([][]int, redN)
		for i := 0; i < redN; i++ {
			input.Scan()
			red[i] = stringArrayToIntArray(strings.Fields(input.Text()))
		}
		input.Scan()
		blueN, _ := strconv.Atoi(input.Text())
		blue := make([][]int, blueN)
		for i := 0; i < blueN; i++ {
			input.Scan()
			blue[i] = stringArrayToIntArray(strings.Fields(input.Text()))
		}
		ans := shortestAlternatingPaths(n, red, blue)
		fmt.Printf("%v\n", ans)
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
 * 执行用时：8ms 在所有Go提交中击败了81.82%的用户
 * 占用内存：5.1MB 在所有Go提交中击败了77.27%的用户
**/
