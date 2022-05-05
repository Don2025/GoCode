package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func canVisitAllRooms(rooms [][]int) bool {
	n := len(rooms)
	visited := make([]bool, n)
	visited[0] = true
	queue := []int{0}
	var num int
	for len(queue) != 0 {
		x := queue[0]
		queue = queue[1:]
		num++
		for _, it := range rooms[x] {
			if !visited[it] {
				visited[it] = true
				queue = append(queue, it)
			}
		}
	}
	return num == n
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		rooms := make([][]int, n)
		for i := 0; i < n; i++ {
			input.Scan()
			rooms[i] = stringArrayToIntArray(strings.Fields(input.Text()))
		}
		println(canVisitAllRooms(rooms))
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
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：3.9MB 在所有Go提交中击败了92.16%的用户
**/
