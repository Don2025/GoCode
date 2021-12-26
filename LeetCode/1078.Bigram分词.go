package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func findOcurrences(text string, first string, second string) []string {
	arr := strings.Fields(text)
	var ans []string
	for i := 0; i < len(arr)-2; i++ {
		if arr[i] == first && arr[i+1] == second {
			ans = append(ans, arr[i+2])
		}
	}
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		text := input.Text()
		input.Scan()
		arr := strings.Fields(input.Text())
		first, second := arr[0], arr[1]
		ans := findOcurrences(text, first, second)
		for _, x := range ans {
			fmt.Printf("%s ", x)
		}
		fmt.Println()
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2MB 在所有Go提交中击败了56.00%的用户
**/
