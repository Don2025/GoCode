package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

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

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		board := make([][]int, n)
		for i := 0; i < n; i++ {
			input.Scan()
			board[i] = stringArrayToIntArray(strings.Fields(input.Text()))
		}
		println(snakesAndLadders(board))
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
 * 执行用时：28ms 在所有Go提交中击败了14.29%的用户
 * 占用内存：6.7MB 在所有Go提交中击败了14.29%的用户
**/
