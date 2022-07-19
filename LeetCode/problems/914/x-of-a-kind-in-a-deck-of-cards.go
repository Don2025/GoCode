package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
)

// https://leetcode.cn/problems/x-of-a-kind-in-a-deck-of-cards/
//------------------------Leetcode Problem 914------------------------
func hasGroupsSizeX(deck []int) bool {
	mp := make(map[int]int)
	for _, x := range deck {
		mp[x]++
	}
	var cnt int
	for _, x := range mp {
		if x > 0 {
			cnt = gcd(cnt, x)
			if cnt == 1 {
				return false
			}
		}
	}
	return cnt >= 2
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

//------------------------Leetcode Problem 914------------------------
/*
 * https://leetcode.cn/problems/x-of-a-kind-in-a-deck-of-cards/
 * 执行用时：12ms 在所有Go提交中击败了92.68%的用户
 * 占用内存：5.6MB 在所有Go提交中击败了65.85%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		deck := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", hasGroupsSizeX(deck))
	}
}
