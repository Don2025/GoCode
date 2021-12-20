package main

import (
	"bufio"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func findRadius(houses []int, heaters []int) int {
	sort.Ints(heaters)
	var ans int
	for _, house := range houses {
		minDis := math.MaxInt32
		j := sort.SearchInts(heaters, house+1) // j是位置大于当前房屋的供暖器的最小下标
		//在Go中 若数组中没有目标元素则sort.SearchInts的返回值是目标元素应该插入的位置
		if j < len(heaters) { //如果右边存在供暖器则计算距离
			minDis = heaters[j] - house
		}
		i := j - 1  // i是位置小于等于当前房屋的供暖器的最大下标
		if i >= 0 { // 如果当前房屋左边存在供暖器 且距离小于右手边供暖器 则取最短距离
			minDis = min(minDis, house-heaters[i])
		}
		//查出当前房屋与供暖器的最短距离后 若半径无法满足当前房屋则让它满足当前房屋
		if minDis > ans {
			ans = minDis
		}
	}
	return ans
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
		houses := stringArrayToIntArray(strings.Fields(input.Text()))
		input.Scan()
		heaters := stringArrayToIntArray(strings.Fields(input.Text()))
		println(findRadius(houses, heaters))
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
 * 执行用时：40ms 在所有Go提交中击败了98.61%的用户
 * 占用内存：6.7MB 在所有Go提交中击败了66.67%的用户
**/
