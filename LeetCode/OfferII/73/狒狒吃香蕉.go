package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"sort"
	"strconv"
)

//  https://leetcode.cn/problems/nZZqjQ/
// ------------------------剑指 Offer II Problem 73------------------------
func minEatingSpeed(piles []int, h int) int {
	var max int
	for _, pile := range piles {
		if pile > max {
			max = pile
		}
	}
	return 1 + sort.Search(max-1, func(speed int) bool {
		speed++
		var time int
		for _, pile := range piles {
			time += (pile + speed - 1) / speed
		}
		return time <= h
	})
}

// ------------------------剑指 Offer II Problem 73------------------------
/*
 * https://leetcode.cn/problems/nZZqjQ/
 * 执行用时：32ms 在所有Go提交中击败了41.18%的用户
 * 占用内存：6.4MB 在所有Go提交中击败了18.30%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	Printf("Input a line of numbers separated by space:")
	for scanner.Scan() {
		piles := utils.StringToInts(scanner.Text())
		Printf("Input a number:")
		scanner.Scan()
		h, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", minEatingSpeed(piles, h))
		Printf("Input a line of numbers separated by space:")
	}
}
