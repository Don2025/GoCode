package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"sort"
)

// https://leetcode.cn/problems/make-two-arrays-equal-by-reversing-sub-arrays/
//------------------------Leetcode Problem 1460------------------------
func canBeEqual(target []int, arr []int) bool {
	sort.Ints(target)
	sort.Ints(arr)
	for i, x := range target {
		if x != arr[i] {
			return false
		}
	}
	return true
}

//------------------------Leetcode Problem 1460------------------------
/*
 * https://leetcode.cn/problems/make-two-arrays-equal-by-reversing-sub-arrays/
 * 执行用时：12ms 在所有Go提交中击败了26.79%的用户
 * 占用内存：3.9MB 在所有Go提交中击败了80.36%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		target := utils.StringToInts(scanner.Text())
		scanner.Scan()
		arr := utils.StringToInts(scanner.Text())
		Printf("Output: %t\n", canBeEqual(target, arr))
	}
}
