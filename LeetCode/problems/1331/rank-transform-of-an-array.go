package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"sort"
)

// https://leetcode.cn/problems/rank-transform-of-an-array/
//------------------------Leetcode Problem 1331------------------------
func arrayRankTransform(arr []int) []int {
	tmp := make([]int, len(arr))
	copy(tmp, arr)
	sort.Ints(tmp)
	rank := make(map[int]int)
	for _, v := range tmp {
		if _, ok := rank[v]; !ok {
			rank[v] = len(rank) + 1
		}
	}
	for k, v := range arr {
		arr[k] = rank[v]
	}
	return arr
}

//------------------------Leetcode Problem 1331------------------------
/*
 * https://leetcode.cn/problems/rank-transform-of-an-array/
 * 执行用时：68ms 在所有Go提交中击败了54.17%的用户
 * 占用内存：12.3MB 在所有Go提交中击败了52.08%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		arr := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", arrayRankTransform(arr))
	}
}
