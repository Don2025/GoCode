package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func merge(intervals [][]int) [][]int {
	var ans [][]int
	if len(intervals) == 0 {
		return ans
	}
	sort.Ints(intervals)

}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		intervals := make([][]int, n)
		for i := range intervals {
			input.Scan()
			intervals[i] = stringArrayToIntArray(strings.Fields(input.Text()))
		}
		ans := merge(intervals)
		for _, row := range ans {
			fmt.Printf("%v ", row)
		}
		fmt.Println()
	}
}
