package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

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

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		hand := stringArrayToIntArray(strings.Fields(input.Text()))
		input.Scan()
		groupSize, _ := strconv.Atoi(input.Text())
		println(isNStraightHand(hand, groupSize))
	}
}

func stringArrayToIntArray(strings []string) []int {
	var arr []int
	for _, x := range strings {
		n, _ := strconv.Atoi(x)
		arr = append(arr, n)
	}
	return arr
}

/*
 * 执行用时：48ms 在所有Go提交中击败了48.57%的用户
 * 占用内存：6.7MB 在所有Go提交中击败了48.57%的用户
**/
