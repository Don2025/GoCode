package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	var ans [][]int
	for i, j := 0, 0; i < len(firstList) && j < len(secondList); {
		l := max(firstList[i][0], secondList[j][0])
		r := min(firstList[i][1], secondList[j][1])
		if l <= r {
			ans = append(ans, []int{l, r})
		}
		if firstList[i][1] < secondList[j][1] {
			i++
		} else {
			j++
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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
		n, _ := strconv.Atoi(input.Text())
		arr1 := make([][]int, n)
		for i := range arr1 {
			input.Scan()
			arr1[i] = stringArrayToIntArray(strings.Fields(input.Text()))
		}
		input.Scan()
		n, _ = strconv.Atoi(input.Text())
		arr2 := make([][]int, n)
		for i := range arr2 {
			input.Scan()
			arr2[i] = stringArrayToIntArray(strings.Fields(input.Text()))
		}
		ans := intervalIntersection(arr1, arr2)
		for _, x := range ans {
			fmt.Printf("%v ", x)
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
 * 执行用时：16ms 在所有Go提交中击败了90.71%的用户
 * 占用内存：6.7MB 在所有Go提交中击败了30.77%的用户
**/
