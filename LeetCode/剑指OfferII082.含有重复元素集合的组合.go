package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var freq [][2]int
	for _, x := range candidates {
		if freq == nil || x != freq[len(freq)-1][0] {
			freq = append(freq, [2]int{x, 1})
		} else {
			freq[len(freq)-1][1]++
		}
	}
	var ans [][]int
	var sequence []int
	var dfs func(int, int)
	dfs = func(pos, rest int) {
		if rest == 0 {
			ans = append(ans, append([]int(nil), sequence...))
			return
		}
		if pos == len(freq) || rest < freq[pos][0] {
			return
		}
		dfs(pos+1, rest)
		most := min(rest/freq[pos][0], freq[pos][1])
		for i := 1; i <= most; i++ {
			sequence = append(sequence, freq[pos][0])
			dfs(pos+1, rest-i*freq[pos][0])
		}
		sequence = sequence[:len(sequence)-most]
	}
	dfs(0, target)
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		candidates := stringArrayToIntArray(strings.Fields(input.Text()))
		input.Scan()
		target, _ := strconv.Atoi(input.Text())
		ans := combinationSum(candidates, target)
		for _, row := range ans {
			fmt.Printf("%v ", row)
		}
		fmt.Println()
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.4MB 在所有Go提交中击败了61.54%的用户
**/
