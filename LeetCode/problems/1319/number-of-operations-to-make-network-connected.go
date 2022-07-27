package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/number-of-operations-to-make-network-connected/
//------------------------Leetcode Problem 1319------------------------
func makeConnected(n int, connections [][]int) int {
	if len(connections) < n-1 { // n个连通结点至少需要n-1条线
		return -1
	}
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	var find func(int) int
	find = func(node int) int {
		if parent[node] == node {
			return node
		}
		return find(parent[node])
	}
	var union func(int, int)
	union = func(node1 int, node2 int) {
		node1 = find(node1)
		node2 = find(node2)
		if node1 == node2 {
			return
		}
		parent[node1] = node2
	}
	for _, con := range connections {
		union(con[0], con[1])
	}
	var cnt int
	for i := range parent {
		if parent[i] == i {
			cnt++
		}
	}
	return cnt - 1
}

//------------------------Leetcode Problem 1319------------------------
/*
 * https://leetcode.cn/problems/number-of-operations-to-make-network-connected/
 * 执行用时：124ms 在所有Go提交中击败了11.11%的用户
 * 占用内存：9.5MB 在所有Go提交中击败了78.89%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		m, _ := strconv.Atoi(scanner.Text())
		connections := make([][]int, m)
		for i := range connections {
			scanner.Scan()
			connections[i] = utils.StringToInts(scanner.Text())
		}
		Printf("Output: %v\n", makeConnected(n, connections))
	}
}
