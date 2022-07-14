package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"math"
	"os"
	"sort"
)

// https://leetcode.cn/problems/minimum-absolute-difference/
//------------------------Leetcode Problem 1200------------------------
func minimumAbsDifference(arr []int) (ans [][]int) {
	sort.Ints(arr)
	for i, best := 0, math.MaxInt32; i < len(arr)-1; i++ {
		if t := arr[i+1] - arr[i]; t < best {
			best = t
			ans = [][]int{{arr[i], arr[i+1]}}
		} else if t == best {
			ans = append(ans, []int{arr[i], arr[i+1]})
		}
	}
	return
}

//------------------------Leetcode Problem 1200------------------------
/*
 * https://leetcode.cn/problems/minimum-absolute-difference/
 * 执行用时：64ms 在所有Go提交中击败了82.19%的用户
 * 占用内存：8.4MB 在所有Go提交中击败了17.81%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		arr := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", minimumAbsDifference(arr))
	}
}
