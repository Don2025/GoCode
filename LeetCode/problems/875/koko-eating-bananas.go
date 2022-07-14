package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"sort"
	"strconv"
)

// https://leetcode.cn/problems/koko-eating-bananas/
//------------------------Leetcode Problem 875------------------------
func minEatingSpeed(piles []int, h int) int {
	var ans int
	for _, pile := range piles {
		if pile > ans {
			ans = pile
		}
	}
	return 1 + sort.Search(ans-1, func(speed int) bool {
		speed++
		time := 0
		for _, pile := range piles {
			time += (pile + speed - 1) / speed
		}
		return time <= h
	})
}

//------------------------Leetcode Problem 875------------------------
/*
 * https://leetcode.cn/problems/koko-eating-bananas/
 * 执行用时：28ms 在所有Go提交中击败了80.04%的用户
 * 占用内存：6.4MB 在所有Go提交中击败了18.43%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		piles := utils.StringToInts(scanner.Text())
		scanner.Scan()
		h, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", minEatingSpeed(piles, h))
	}
}
