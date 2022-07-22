package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"sort"
)

// https://leetcode.cn/problems/array-of-doubled-pairs/
//------------------------Leetcode Problem 954------------------------
func canReorderDoubled(arr []int) bool {
	m := make(map[int]int)
	for _, x := range arr {
		m[x]++
	}
	if m[0]&1 == 1 {
		return false
	}
	v := make([]int, 0, len(m))
	for x := range m {
		v = append(v, x)
	}
	sort.Slice(v, func(i, j int) bool { return abs(v[i]) < abs(v[j]) })
	for _, x := range v {
		if m[x<<1] < m[x] {
			return false
		}
		m[x*2] -= m[x]
	}
	return true
}

func abs(n int) int {
	if n < 0 {
		n = -n
	}
	return n
}

//------------------------Leetcode Problem 954------------------------
/*
 * https://leetcode.cn/problems/array-of-doubled-pairs/
 * 执行用时：68ms 在所有Go提交中击败了90.00%的用户
 * 占用内存：7MB 在所有Go提交中击败了83.33%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		arr := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", canReorderDoubled(arr))
	}
}
