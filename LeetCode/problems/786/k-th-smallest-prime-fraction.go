package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"sort"
	"strconv"
)

// https://leetcode.cn/problems/k-th-smallest-prime-fraction/
//------------------------Leetcode Problem 786------------------------
func kthSmallestPrimeFraction(arr []int, k int) []int {
	type pair struct {
		x, y int
	}
	frac := make([]pair, 0, len(arr)*(len(arr)-1)/2)
	for i, x := range arr {
		for _, y := range arr[i+1:] {
			frac = append(frac, pair{x, y})
		}
	}
	sort.Slice(frac, func(i, j int) bool {
		return frac[i].x*frac[j].y < frac[i].y*frac[j].x
	})
	return []int{frac[k-1].x, frac[k-1].y}
}

//------------------------Leetcode Problem 786------------------------
/*
 * https://leetcode.cn/problems/k-th-smallest-prime-fraction/
 * 执行用时：604ms 在所有Go提交中击败了14.29%的用户
 * 占用内存：23.9MB 在所有Go提交中击败了14.29%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		arr := utils.StringToInts(scanner.Text())
		scanner.Scan()
		k, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", kthSmallestPrimeFraction(arr, k))
	}
}
