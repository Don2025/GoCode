package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findClosestElements(arr []int, k int, x int) []int {
	l, r := 0, len(arr)-1-k
	for l <= r {
		mid := l + (r-l)>>1
		if x-arr[mid] > arr[mid+k]-x {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return arr[l : l+k]
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		arr := stringArrayToIntArray(strings.Fields(input.Text()))
		input.Scan()
		k, _ := strconv.Atoi(input.Text())
		input.Scan()
		x, _ := strconv.Atoi(input.Text())
		ans := findClosestElements(arr, k, x)
		fmt.Printf("%v\n", ans)
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
 * 执行用时：40ms 在所有Go提交中击败了34.01%的用户
 * 占用内存：6.8MB 在所有Go提交中击败了73.28%的用户
**/
