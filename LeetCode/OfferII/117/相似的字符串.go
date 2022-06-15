package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
)

// https://leetcode.cn/problems/H6lPxb/
// ------------------------剑指 Offer II Problem 117------------------------
type unionFind struct {
	Parent, size []int
	SetCount     int
}

func newUnionFind(n int) *unionFind {
	parent := make([]int, n)
	size := make([]int, n)
	for i := range parent {
		parent[i] = i
		size[i] = 1
	}
	return &unionFind{parent, size, n}
}

func (uf *unionFind) find(x int) int {
	if uf.Parent[x] != x {
		uf.Parent[x] = uf.find(uf.Parent[x])
	}
	return uf.Parent[x]
}

func (uf *unionFind) union(x, y int) {
	fx, fy := uf.find(x), uf.find(y)
	if fx == fy {
		return
	}
	if uf.size[fx] < uf.size[fy] {
		fx, fy = fy, fx
	}
	uf.size[fx] += uf.size[fy]
	uf.Parent[fy] = fx
	uf.SetCount--
}

func (uf *unionFind) inSameSet(x, y int) bool {
	return uf.find(x) == uf.find(y)
}

func isSimilar(s, t string) bool {
	diff := 0
	for i := range s {
		if s[i] != t[i] {
			diff++
			if diff > 2 {
				return false
			}
		}
	}
	return true
}

func numSimilarGroups(strs []string) int {
	n := len(strs)
	uf := newUnionFind(n)
	for i, s := range strs {
		for j := i + 1; j < n; j++ {
			if !uf.inSameSet(i, j) && isSimilar(s, strs[j]) {
				uf.union(i, j)
			}
		}
	}
	return uf.SetCount
}

// ------------------------剑指 Offer II Problem 117------------------------
/*
 * https://leetcode.cn/problems/H6lPxb/
 * 执行用时：8ms 在所有Go提交中击败了79.71%的用户
 * 占用内存：4MB 在所有Go提交中击败了47.83%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	Printf("Input a line of strings separated by space:")
	for scanner.Scan() {
		strs := strings.Fields(scanner.Text())
		Printf("Output: %v\n", numSimilarGroups(strs))
		Printf("Input a line of strings separated by space:")
	}
}
