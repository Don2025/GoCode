package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"sort"
	"strconv"
)

// https://leetcode.cn/problems/relative-ranks/
//------------------------Leetcode Problem 506------------------------
func findRelativeRanks(score []int) []string {
	type pair struct {
		score, rank int
	}
	arr := make([]pair, len(score))
	for i, x := range score {
		arr[i] = pair{x, i}
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].score > arr[j].score
	})
	model := [3]string{"Gold Medal", "Silver Medal", "Bronze Medal"}
	ans := make([]string, len(score))
	for i, p := range arr {
		if i < 3 {
			ans[p.rank] = model[i]
		} else {
			ans[p.rank] = strconv.Itoa(i + 1)
		}
	}
	return ans
}

//------------------------Leetcode Problem 506------------------------
/*
 * https://leetcode.cn/problems/relative-ranks/
 * 执行用时：8ms 在所有Go提交中击败了99.01%的用户
 * 占用内存：5.2MB 在所有Go提交中击败了79.21%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		score := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", findRelativeRanks(score))
	}
}
