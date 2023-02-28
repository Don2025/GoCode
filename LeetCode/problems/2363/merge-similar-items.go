package main

import (
	"bufio"
	"fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"sort"
	"strconv"
)

// https://leetcode.cn/problems/merge-similar-items/
//------------------------Leetcode Problem 2363------------------------
func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
	m := map[int]int{}
	for _, x := range items1 {
		m[x[0]] += x[1]
	}
	for _, x := range items2 {
		m[x[0]] += x[1]
	}
	var ans [][]int
	for k, v := range m {
		ans = append(ans, []int{k, v})
	}
	sort.Slice(ans, func(i, j int) bool {
		return ans[i][0] < ans[j][0]
	})
	return ans
}

//------------------------Leetcode Problem 2363------------------------
/*
 * https://leetcode.cn/problems/merge-similar-items/
 * 执行用时：20ms 在所有Go提交中击败了20.00%的用户
 * 占用内存：6.6MB 在所有Go提交中击败了11.43%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		items1 := make([][]int, n)
		for i := range items1 {
			scanner.Scan()
			items1[i] = utils.StringToInts(scanner.Text())
		}
		scanner.Scan()
		n, _ = strconv.Atoi(scanner.Text())
		items2 := make([][]int, n)
		for i := range items2 {
			scanner.Scan()
			items2[i] = utils.StringToInts(scanner.Text())
		}
		ans := mergeSimilarItems(items1, items2)
		fmt.Printf("%d\n", ans)
	}
}
