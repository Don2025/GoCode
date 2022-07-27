package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"sort"
	"strconv"
)

// https://leetcode.cn/problems/maximum-ice-cream-bars/
//------------------------Leetcode Problem 1833------------------------
func maxIceCream(costs []int, coins int) int {
	sort.Ints(costs)
	var cnt int
	for _, x := range costs {
		if coins >= x {
			coins -= x
		} else {
			break
		}
		cnt++
	}
	return cnt
}

//------------------------Leetcode Problem 1833------------------------
/*
 * https://leetcode.cn/problems/maximum-ice-cream-bars/
 * 执行用时：232ms 在所有Go提交中击败了50.45%的用户
 * 占用内存：9.2MB 在所有Go提交中击败了82.28%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		costs := utils.StringToInts(scanner.Text())
		scanner.Scan()
		coins, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", maxIceCream(costs, coins))
	}
}
