package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	if d < 0 {
		d = -d
	}
	sort.Ints(arr1)
	sort.Ints(arr2)
	var ans, l, r int
	for i := range arr1 {
		l = l + sort.Search(len(arr2[l:]), func(index int) bool {
			return arr2[l+index] >= arr1[i]-d
		})
		r = r + sort.Search(len(arr2[r:]), func(index int) bool {
			return arr2[r+index] > arr1[i]+d
		})
		if l >= r {
			ans++
		}
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		arr1 := stringArrayToIntArray(strings.Fields(input.Text()))
		input.Scan()
		arr2 := stringArrayToIntArray(strings.Fields(input.Text()))
		input.Scan()
		d, _ := strconv.Atoi(input.Text())
		println(findTheDistanceValue(arr1, arr2, d))
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
 * 执行用时：8ms 在所有Go提交中击败了54.81%的用户
 * 占用内存：3.6MB 在所有Go提交中击败了12.50%的用户
**/
