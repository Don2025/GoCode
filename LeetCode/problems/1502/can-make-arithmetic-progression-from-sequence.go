package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"sort"
)

// https://leetcode.cn/problems/can-make-arithmetic-progression-from-sequence/
//------------------------Leetcode Problem 1502------------------------
func canMakeArithmeticProgression(arr []int) bool {
	if len(arr) <= 2 {
		return true
	}
	sort.Ints(arr)
	d := arr[1] - arr[0]
	for i := 1; i < len(arr)-1; i++ {
		if arr[i+1]-arr[i] != d {
			return false
		}
	}
	return true
}

//------------------------Leetcode Problem 1502------------------------
/*
 * https://leetcode.cn/problems/can-make-arithmetic-progression-from-sequence/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.4MB 在所有Go提交中击败了95.94%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		arr := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", canMakeArithmeticProgression(arr))
	}
}
