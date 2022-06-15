package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
	"strings"
)

// https://leetcode.cn/problems/vlzXQL/
// ------------------------剑指 Offer II Problem 111------------------------
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	// 给方程组中的每个变量编号
	id := map[string]int{}
	for _, eq := range equations {
		a, b := eq[0], eq[1]
		if _, has := id[a]; !has {
			id[a] = len(id)
		}
		if _, has := id[b]; !has {
			id[b] = len(id)
		}
	}

	fa := make([]int, len(id))
	w := make([]float64, len(id))
	for i := range fa {
		fa[i] = i
		w[i] = 1
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			f := find(fa[x])
			w[x] *= w[fa[x]]
			fa[x] = f
		}
		return fa[x]
	}
	merge := func(from, to int, val float64) {
		fFrom, fTo := find(from), find(to)
		w[fFrom] = val * w[to] / w[from]
		fa[fFrom] = fTo
	}

	for i, eq := range equations {
		merge(id[eq[0]], id[eq[1]], values[i])
	}

	ans := make([]float64, len(queries))
	for i, q := range queries {
		start, hasS := id[q[0]]
		end, hasE := id[q[1]]
		if hasS && hasE && find(start) == find(end) {
			ans[i] = w[start] / w[end]
		} else {
			ans[i] = -1
		}
	}
	return ans
}

// ------------------------剑指 Offer II Problem 111------------------------
/*
 * https://leetcode.cn/problems/vlzXQL/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了91.67%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	Printf("Input the size of [][]string:")
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		equations := make([][]string, n)
		for i := range equations {
			Printf("Input a line of strings separated by space:")
			scanner.Scan()
			equations[i] = strings.Fields(scanner.Text())
		}
		Printf("Input a line of floating point numbers separated by space:")
		scanner.Scan()
		values := utils.StringToFloat64s(scanner.Text())
		Printf("Input the size of [][]string:")
		scanner.Scan()
		n, _ = strconv.Atoi(scanner.Text())
		queries := make([][]string, n)
		for i := range queries {
			Printf("Input a line of strings separated by space:")
			scanner.Scan()
			queries[i] = strings.Fields(scanner.Text())
		}
		Printf("Output: %v\n", calcEquation(equations, values, queries))
		Printf("Input a line of strings separated by space:")
	}
}
