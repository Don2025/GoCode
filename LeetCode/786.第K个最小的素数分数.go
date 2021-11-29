package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func kthSmallestPrimeFraction(arr []int, k int) []int {
	type pair struct {
		x, y int
	}
	frac := make([]pair, 0, len(arr)*(len(arr)-1)/2)
	for i, x := range arr {
		for _, y := range arr[i+1:] {
			frac = append(frac, pair{x, y})
		}
	}
	sort.Slice(frac, func(i, j int) bool {
		return frac[i].x*frac[j].y < frac[i].y*frac[j].x
	})
	return []int{frac[k-1].x, frac[k-1].y}
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		arr := stringArrayToIntArray(strings.Fields(input.Text()))
		input.Scan()
		k, _ := strconv.Atoi(input.Text())
		ans := kthSmallestPrimeFraction(arr, k)
		for _, x := range ans {
			fmt.Printf("%d ", x)
		}
		fmt.Println()
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
 * 执行用时：604ms 在所有Go提交中击败了14.29%的用户
 * 占用内存：23.9MB 在所有Go提交中击败了14.29%的用户
**/
