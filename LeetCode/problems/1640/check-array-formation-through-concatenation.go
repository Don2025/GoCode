package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/check-array-formation-through-concatenation/
//------------------------Leetcode Problem 1640------------------------
func canFormArray(arr []int, pieces [][]int) bool {
	m := map[int]int{}
	for i, p := range pieces {
		m[p[0]] = i
	}
	for i := 0; i < len(arr); {
		j, ok := m[arr[i]]
		if !ok {
			return false
		}
		for _, x := range pieces[j] {
			if x != arr[i] {
				return false
			}
			i++
		}
	}
	return true
}

// ------------------------Leetcode Problem 1640------------------------
/*
 * https://leetcode.cn/problems/check-array-formation-through-concatenation/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.6MB 在所有Go提交中击败了63.64%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		arr := utils.StringToInts(scanner.Text())
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())
		pieces := make([][]int, n)
		for i := range pieces {
			pieces[i] = utils.StringToInts(scanner.Text())
		}
		Printf("Output: %t\n", canFormArray(arr, pieces))
	}
}
