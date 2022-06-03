package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func combine(n int, k int) [][]int {
	var ans [][]int
	var tmp []int
	for i := 1; i <= k; i++ {
		tmp = append(tmp, i)
	}
	tmp = append(tmp, n+1)
	for i := 0; i < k; {
		a := make([]int, k)
		copy(a, tmp[:k])
		ans = append(ans, a)
		for i = 0; i < k && tmp[i]+1 == tmp[i+1]; i++ {
			tmp[i] = i + 1
		}
		tmp[i]++
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		arr := strings.Fields(input.Text())
		n, _ := strconv.Atoi(arr[0])
		k, _ := strconv.Atoi(arr[1])
		ans := combine(n, k)
		for _, row := range ans {
			fmt.Printf("%v ", row)
		}
		fmt.Println()
	}
}

/*
 * 执行用时：8ms 在所有Go提交中击败了70.80%的用户
 * 占用内存：6.2MB 在所有Go提交中击败了94.69%的用户
**/
