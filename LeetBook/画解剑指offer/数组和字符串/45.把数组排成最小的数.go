package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

//输入一个非负整数数组，把数组里所有数字拼接起来排成一个数，打印能拼接出的所有数字中最小的一个。
func minNumber(nums []int) string {
	var strArray []string
	for _, x := range nums {
		s := strconv.Itoa(x)
		strArray = append(strArray, s)
	}
	sort.Slice(strArray, func(i, j int) bool {
		return strArray[i]+strArray[j] < strArray[j]+strArray[i]
	})
	var ans string
	for _, str := range strArray {
		ans += str
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		nums := stringArrayToIntArray(strings.Fields(input.Text()))
		println(minNumber(nums))
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
 * 执行用时：4ms 在所有Go提交中击败了84.86%的用户
 * 占用内存：2.9MB 在所有Go提交中击败了29.22%的用户
**/
