package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/keys-and-rooms/
//------------------------Leetcode Problem 841------------------------
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

//------------------------Leetcode Problem 841------------------------
/*
 * https://leetcode.cn/problems/keys-and-rooms/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：3.9MB 在所有Go提交中击败了92.16%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		rooms := make([][]int, n)
		for i := 0; i < n; i++ {
			scanner.Scan()
			rooms[i] = utils.StringToInts(scanner.Text())
		}
		Printf("Output: %v\n", canVisitAllRooms(rooms))
	}
}
