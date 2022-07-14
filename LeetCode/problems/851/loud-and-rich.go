package main

import (
	"bufio"
	. "fmt"
	"github.com/Don2025/GoCode/utils"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/loud-and-rich/
//------------------------Leetcode Problem 851------------------------
func loudAndRich(richer [][]int, quiet []int) []int {
	g := make([][]int, len(quiet)) //有向无环图，有钱人指向没钱人
	in := make([]int, len(quiet))  //记录没钱人的入度，即比他有钱的人数
	for _, r := range richer {
		g[r[0]] = append(g[r[0]], r[1])
		in[r[1]]++
	}
	ans := make([]int, len(quiet)) //最有钱的ans是他自己
	for i := range ans {
		ans[i] = i
	}
	//广度优先遍历,先把最有钱的人入队，即入度为0的
	queue := make([]int, 0, len(quiet))
	for i, x := range in {
		if x == 0 {
			queue = append(queue, i)
		}
	}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for _, x := range g[cur] {
			if quiet[ans[cur]] < quiet[ans[x]] {
				ans[x] = ans[cur]
			}
			in[x]--
			if in[x] == 0 {
				queue = append(queue, x)
			}
		}
	}
	return ans
}

//------------------------Leetcode Problem 851------------------------
/*
 * https://leetcode.cn/problems/loud-and-rich/
 * 执行用时：76ms 在所有Go提交中击败了5.56%的用户
 * 占用内存：8.2MB 在所有Go提交中击败了27.78%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		richer := make([][]int, n)
		for i := 0; i < n; i++ {
			scanner.Scan()
			richer[i] = utils.StringToInts(scanner.Text())
		}
		scanner.Scan()
		quiet := utils.StringToInts(scanner.Text())
		Printf("Output: %v\n", loudAndRich(richer, quiet))
	}
}
