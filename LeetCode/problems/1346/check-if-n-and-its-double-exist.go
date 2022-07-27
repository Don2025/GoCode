package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/check-if-n-and-its-double-exist/
//------------------------Leetcode Problem 1346------------------------
func checkIfExist(arr []int) bool {
	m := make(map[int]int, len(arr))
	for k, v := range arr {
		m[v] = k
	}
	for i := range arr {
		arr[i] <<= 1
	}
	for i, x := range arr {
		if j, ok := m[x]; i != j && ok {
			return true
		}
	}
	return false
}

//------------------------Leetcode Problem 1346------------------------
/*
 * https://leetcode.cn/problems/check-if-n-and-its-double-exist/
 * 执行用时：4ms 在所有Go提交中击败了80s.99%的用户
 * 占用内存：3.6MB 在所有Go提交中击败了23.14%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		arr := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", checkIfExist(arr))
	}
}
