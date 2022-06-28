package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"sort"
	"strconv"
)

// https://leetcode.cn/problems/top-k-frequent-elements/
//------------------------Leetcode Problem 347------------------------
func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, x := range nums {
		m[x]++
	}
	var ans []int
	for x := range m {
		ans = append(ans, x)
	}
	sort.Slice(ans, func(i, j int) bool {
		return m[ans[i]] > m[ans[j]]
	})
	return ans[:k]
}

//------------------------Leetcode Problem 347------------------------
/*
 * https://leetcode.cn/problems/top-k-frequent-elements/
 * 执行用时：12ms 在所有Go提交中击败了82.90%的用户
 * 占用内存：5.3MB 在所有Go提交中击败了86.96%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums := utils.StringToInts(scanner.Text())
		scanner.Scan()
		k, _ := strconv.Atoi(scanner.Text())
		Printf("%v\n", topKFrequent(nums, k))
	}
}
