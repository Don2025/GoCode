package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func combine(n int, k int) [][]int {
	var ans [][]int
	var t []int
	var dfs func(int)
	dfs = func(i int) {
		if len(t)+(n-i+1) < k {
			return
		}
		if len(t) == k {
			arr := make([]int, k)
			copy(arr, t)
			ans = append(ans, arr)
			return
		}
		for j := i; j <= n; j++ {
			t = append(t, j)
			dfs(j + 1)
			t = t[:len(t)-1]
		}
	}
	dfs(1)
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		input.Scan()
		k, _ := strconv.Atoi(input.Text())
		ans := combine(n, k)
		for _, x := range ans {
			fmt.Printf("%d ", x)
		}
		fmt.Println()
	}
}

/*
 * 执行用时：4ms 在所有Go提交中击败了98.57%的用户
 * 占用内存：6.3MB 在所有Go提交中击败了83.33%的用户
**/
