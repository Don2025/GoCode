package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"sort"
)

// https://leetcode.cn/problems/mean-of-array-after-removing-some-elements/
//------------------------Leetcode Problem 670------------------------
func trimMean(arr []int) float64 {
	sort.Ints(arr)
	n := len(arr)
	ans := 0
	for _, x := range arr[n/20 : 19*n/20] {
		ans += x
	}
	return float64(ans*10) / float64(n*9)
}

//------------------------Leetcode Problem 670------------------------
/*
 * https://leetcode.cn/problems/mean-of-array-after-removing-some-elements/
 * 执行用时：4ms 在所有Go提交中击败了94.87%的用户
 * 占用内存：3.2MB 在所有Go提交中击败了51.28%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		arr := utils.StringToInts(scanner.Text())
		Printf("Output: %f\n", trimMean(arr))
	}
}
