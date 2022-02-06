package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func minReorder(n int, connections [][]int) int {
	var ans int
	visited := make([]bool, len(connections))
	m := make(map[int][]int, 0)
	for i, c := range connections {
		m[c[0]] = append(m[c[0]], i)
		m[c[1]] = append(m[c[1]], i)
	}
	q := []int{0}
	for len(q) != 0 {
		t := q[0]
		q = q[1:]
		for _, x := range m[t] {
			if !visited[x] {
				visited[x] = true
				if connections[x][1] == t {
					q = append(q, connections[x][0])
				} else if connections[x][0] == t {
					q = append(q, connections[x][1])
					ans++
				}
			}
		}
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		input.Scan()
		m, _ := strconv.Atoi(input.Text())
		connections := make([][]int, m)
		for i := range connections {
			input.Scan()
			connections[i] = stringArrayToIntArray(strings.Fields(input.Text()))
		}
		println(minReorder(n, connections))
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
 * 执行用时：204ms 在所有Go提交中击败了50.00%的用户
 * 占用内存：17.5MB 在所有Go提交中击败了39.29%的用户
**/
