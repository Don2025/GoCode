package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

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

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		piles := stringArrayToIntArray(strings.Fields(input.Text()))
		input.Scan()
		h, _ := strconv.Atoi(input.Text())
		println(minEatingSpeed(piles, h))
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
 * 执行用时：28ms 在所有Go提交中击败了80.04%的用户
 * 占用内存：6.4MB 在所有Go提交中击败了18.43%的用户
**/
