package main

import "strings"

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */

type MapSum map[string]int

func Constructor() MapSum {
	return MapSum{}
}

func (m MapSum) Insert(key string, val int) {
	m[key] = val
}

func (m MapSum) Sum(prefix string) int {
	var sum int
	for s, v := range m {
		if strings.HasPrefix(s, prefix) {
			sum += v
		}
	}
	return sum
}

/*
 * 执行用时：0ms 在所有Go提交中击败了100.00%的用户
 * 占用内存：2.5MB 在所有Go提交中击败了84.75%的用户
**/
