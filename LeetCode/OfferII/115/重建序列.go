package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/ur2n8P/
// ------------------------剑指 Offer II Problem 115------------------------
func sequenceReconstruction(org []int, seqs [][]int) bool {
	graph := make([][]int, len(org)+1)
	inCount := make([]int, len(org)+1)
	visited := make([]bool, len(org)+1)
	max := len(org)
	for _, arr := range seqs {
		for i, v := range arr {
			if v > max || v < 0 {
				return false
			}
			visited[v] = true
			if i == 0 {
				continue
			}
			graph[arr[i-1]] = append(graph[arr[i-1]], v)
			inCount[v]++
		}
	}
	for i, v := range visited {
		if i == 0 {
			continue
		}
		if !v {
			return false
		}
		visited[i] = false
	}
	for i, c := range inCount {
		if i == 0 {
			continue
		}
		if i != org[0] && c == 0 {
			return false
		}
	}
	for i, from := range org {
		if inCount[from] != 0 {
			return false
		}
		to := -1
		found := false
		if i < len(org)-1 {
			to = org[i+1]
		}
		for _, next := range graph[from] {
			if next == to {
				found = true
			}
			inCount[next]--
		}
		if to > 0 && !found {
			return false
		}
	}
	return true
}

// ------------------------剑指 Offer II Problem 115------------------------
/*
 * https://leetcode.cn/problems/ur2n8P/
 * 执行用时：44ms 在所有Go提交中击败了94.00%的用户
 * 占用内存：8.2MB 在所有Go提交中击败了58.00%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	Printf("Input a line of numbers separated by space:")
	for scanner.Scan() {
		org := utils.StringToInts(scanner.Text())
		Printf("Input the number of rows of matrix:")
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())
		seqs := make([][]int, n)
		for i := range seqs {
			Printf("Input the %d row of numbers separated by spaces:\n", i+1)
			scanner.Scan()
			seqs[i] = utils.StringToInts(scanner.Text())
		}
		Printf("Output: %v\n", sequenceReconstruction(org, seqs))
		Printf("Input a line of numbers separated by space:")
	}
}
