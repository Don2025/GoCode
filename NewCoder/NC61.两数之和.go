package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
 *
 * @param numbers int整型一维数组
 * @param target int整型
 * @return int整型一维数组
 */
func twoSum(numbers []int, target int) []int {
	m := make(map[int]int, len(numbers)) //记录numbers的原始下标+1
	for i, x := range numbers {
		y := target - x
		if j, ok := m[y]; ok {
			return []int{j + 1, i + 1}
		}
		m[x] = i
	}
	return []int{}
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		nums := stringArrayToIntArray(strings.Fields(input.Text()))
		input.Scan()
		target, _ := strconv.Atoi(input.Text())
		ans := twoSum(nums, target)
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
 * 运行时间：4ms 超过17.85%用Go提交的代码
 * 占用内存：1104KB 超过45.13%用Go提交的代码
**/
