package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/7LpjUW/
// ------------------------剑指 Offer II Problem 118------------------------
func findRedundantConnection(edges [][]int) []int {
	parent := make([]int, len(edges)+1)
	for i := range parent {
		parent[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	union := func(from, to int) bool {
		x, y := find(from), find(to)
		if x == y {
			return false
		}
		parent[x] = y
		return true
	}
	for _, e := range edges {
		if !union(e[0], e[1]) {
			return e
		}
	}
	return nil
}

// ------------------------剑指 Offer II Problem 118------------------------
/*
 * https://leetcode.cn/problems/7LpjUW/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.9MB 在所有Go提交中击败了57.65%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	Printf("Input the number of rows of matrix:")
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		edges := make([][]int, n)
		for i := range edges {
			Printf("Input the %d row of numbers separated by spaces:\n", i+1)
			scanner.Scan()
			edges[i] = utils.StringToInts(scanner.Text())
		}
		Printf("Output: %v\n", findRedundantConnection(edges))
		Printf("Input the number of rows of matrix:")
	}
}
