package main

import (
	"bufio"
	. "fmt"
	"math"
	"math/bits"
	"os"
	"strconv"
)

// https://leetcode.cn/problems/smallest-good-base/
//------------------------Leetcode Problem 483------------------------
func smallestGoodBase(n string) string {
	nVal, _ := strconv.Atoi(n)
	mMax := bits.Len(uint(nVal)) - 1 //最长位数是二进制的位数长度
	for m := mMax; m > 1; m-- {      //m是转换为k进制后所有位中1的个数
		l, r := 2, int(math.Pow(float64(nVal), 1/float64(m)))+1
		for l < r { //二分搜索对应的k, k进制
			k := l + (r-l)/2
			sum := 1
			for i := 0; i < m; i++ { //等比数列求和
				sum = sum*k + 1
			}
			if sum < nVal {
				l = k + 1
			} else if sum > nVal {
				r = k - 1
			} else {
				return strconv.Itoa(k)
			}
		}
	}
	return strconv.Itoa(nVal - 1)
}

//------------------------Leetcode Problem 483------------------------
/*
 * https://leetcode.cn/problems/smallest-good-base/
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：1.9MB 在所有Go提交中击败了45.45%的用户
**/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		Printf("Output: %v\n", smallestGoodBase(scanner.Text()))
	}
}
