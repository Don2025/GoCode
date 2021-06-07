package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MaxQueue struct {
	q, max []int
}

func Constructor() MaxQueue {
	return MaxQueue{
		make([]int, 0),
		make([]int, 0),
	}
}

func (this *MaxQueue) Max_value() int {
	if len(this.q) == 0 {
		return -1
	}
	return this.max[0]
}

func (this *MaxQueue) Push_back(value int) {
	this.q = append(this.q, value)
	for len(this.max) != 0 && value > this.max[len(this.max)-1] {
		this.max = this.max[:len(this.max)-1]
	}
	this.max = append(this.max, value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.q) == 0 {
		return -1
	}
	if this.q[0] == this.max[0] {
		this.max = this.max[1:]
	}
	value := this.q[0]
	this.q = this.q[1:]
	return value
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		strArray := strings.Fields(input.Text())
		var queue MaxQueue
		var ans []string
		for _, str := range strArray {
			if str == "MaxQueue" {
				queue = Constructor()
				ans = append(ans, "null")
			} else if str == "push_back" {
				var value int
				print("input a value:")
				fmt.Scanf("%d", &value)
				queue.Push_back(value)
				ans = append(ans, "null")
			} else if str == "pop_front" {
				x := queue.Pop_front()
				ans = append(ans, strconv.Itoa(x))
			} else if str == "max_value" {
				ans = append(ans, strconv.Itoa(queue.Max_value()))
			}
		}
		for _, x := range ans {
			fmt.Printf("%s ", x)
		}
		println()
	}
}

/*
 * 执行用时：104ms 在所有Go提交中击败了58.89%的用户
 * 占用内存：8.5MB 在所有Go提交中击败了30.89%的用户
**/
