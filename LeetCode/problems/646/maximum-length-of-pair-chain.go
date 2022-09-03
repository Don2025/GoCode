package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"sort"
	"strconv"
)

// https://leetcode.cn/problems/maximum-length-of-pair-chain/
//------------------------Leetcode Problem 646------------------------
func findLongestChain(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool { return pairs[i][1] < pairs[j][1] })
	cur := pairs[0][1]
	ans := 1
	for i := 1; i < len(pairs); i++ {
		if cur < pairs[i][0] {
			cur = pairs[i][1]
			ans++
		}
	}
	return ans
}

//------------------------Leetcode Problem 646------------------------
/*
 * https://leetcode.cn/problems/maximum-length-of-pair-chain/
 * 执行用时：28ms 在所有Go提交中击败了89.19%的用户
 * 占用内存：6.5MB 在所有Go提交中击败了83.78%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		pairs := make([][]int, n)
		for i := range pairs {
			scanner.Scan()
			pairs[i] = utils.StringToInts(scanner.Text())
		}
		Printf("Output: %v\n", findLongestChain(pairs))
	}
}
