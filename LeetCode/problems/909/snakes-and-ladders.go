package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/snakes-and-ladders/
//------------------------Leetcode Problem 909------------------------
func snakesAndLadders(board [][]int) int {
	arr := []int{0: -1}
	var cnt, num int
	for i := len(board) - 1; i >= 0; i-- {
		if cnt&1 == 0 {
			for j := 0; j < len(board[i]); j++ {
				arr = append(arr, board[i][j])
			}
		} else {
			for j := len(board[i]) - 1; j >= 0; j-- {
				arr = append(arr, board[i][j])
			}
		}
		cnt++
	}
	target := len(board) * len(board[0])
	type pair struct {
		First, Second int
	}
	visited := make(map[int]bool)
	visited[1] = true
	queue := []pair{0: {1, 0}}
	for len(queue) != 0 {
		num, cnt = queue[0].First, queue[0].Second
		queue = queue[1:]
		if num == target {
			return cnt
		}
		for i := 1; i <= 6 && num+i <= target; i++ {
			if arr[num+i] == -1 && !visited[num+i] {
				visited[num+i] = true
				queue = append(queue, pair{num + i, cnt + 1})
			} else if arr[num+i] != -1 && !visited[arr[num+i]] {
				visited[arr[num+i]] = true
				queue = append(queue, pair{arr[num+i], cnt + 1})
			}
		}
	}
	return -1
}

//------------------------Leetcode Problem 909------------------------
/*
 * https://leetcode.cn/problems/snakes-and-ladders/
 * 执行用时：28ms 在所有Go提交中击败了14.29%的用户
 * 占用内存：6.7MB 在所有Go提交中击败了14.29%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		board := make([][]int, n)
		for i := 0; i < n; i++ {
			scanner.Scan()
			board[i] = utils.StringToInts(scanner.Text())
		}
		Printf("Output: %v\n", snakesAndLadders(board))
	}
}
