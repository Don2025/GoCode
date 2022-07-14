package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"sort"
	"strconv"
)

// https://leetcode.cn/problems/hand-of-straights/
//------------------------Leetcode Problem 846------------------------
func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false
	}
	cnt := make(map[int]int)
	for _, x := range hand {
		cnt[x]++
	}
	sort.Ints(hand)
	for i := 0; i < len(hand); i++ {
		for cnt[hand[i]] > 0 {
			for j := hand[i]; j < hand[i]+groupSize; j++ {
				if cnt[j] == 0 {
					return false
				}
				cnt[j]--
			}
		}
	}
	return true
}

//------------------------Leetcode Problem 846------------------------
/*
 * https://leetcode.cn/problems/hand-of-straights/
 * 执行用时：48ms 在所有Go提交中击败了48.57%的用户
 * 占用内存：6.7MB 在所有Go提交中击败了48.57%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		hand := utils.StringToInts(scanner.Text())
		scanner.Scan()
		groupSize, _ := strconv.Atoi(scanner.Text())
		Printf("Output: %v\n", isNStraightHand(hand, groupSize))
	}
}
