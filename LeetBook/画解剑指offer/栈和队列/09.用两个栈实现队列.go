package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Go里面没有stack，所以没有用两个栈来模拟，直接用container/list来做的
type CQueue struct {
	list *list.List
}

func Constructor() CQueue {
	return CQueue{list.New()}
}

func (this *CQueue) AppendTail(value int) {
	this.list.PushBack(value)
}

func (this *CQueue) DeleteHead() int {
	it := this.list.Front()
	if it != nil {
		this.list.Remove(it)
		return it.Value.(int)
	}
	return -1
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		strArray := strings.Fields(input.Text())
		var queue CQueue
		var ans []string
		for _, str := range strArray {
			if str == "CQueue" {
				queue = Constructor()
				ans = append(ans, "null")
			} else if str == "appendTail" {
				var value int
				print("input a value:")
				fmt.Scanf("%d", &value)
				queue.AppendTail(value)
				ans = append(ans, "null")
			} else if str == "deleteHead" {
				x := queue.DeleteHead()
				ans = append(ans, strconv.Itoa(x))
			}
		}
		for _, x := range ans {
			fmt.Printf("%s ", x)
		}
		println()
	}
}

/*
 * 执行用时：248ms 在所有Go提交中击败了55.83%的用户
 * 占用内存：8.7MB 在所有Go提交中击败了11.40%的用户
**/
