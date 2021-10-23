package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func constructRectangle(area int) []int {
	w := int(math.Sqrt(float64(area)))
	for area%w > 0 {
		w--
	}
	return []int{area / w, w}
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		n, _ := strconv.Atoi(input.Text())
		ans := constructRectangle(n)
		for _, x := range ans {
			fmt.Printf("%d ", x)
		}
		fmt.Println()
	}
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了67.90%的用户
**/
