package main

import (
	"container/ring"
	"fmt"
)

func main() {
	const rLen = 5
	// 创建一个长度为rLen的环形链表
	r := ring.New(rLen)
	fmt.Printf("r.Len() is %d.\n", r.Len())
	// 用几个整数初始化环元素的值
	for i := 0; i < rLen; i++ {
		r.Value = i
		r = r.Next()
	}
	// 从前往后遍历环形链表并打印环元素的值
	for i := 0; i < rLen; i++ {
		fmt.Printf("%d ", r.Value)
		r = r.Next()
	}
	// 0 1 2 3 4
	println()
	// 从后往前遍历环形链表并打印环元素的值
	for i := 0; i < rLen; i++ {
		r = r.Prev()
		fmt.Printf("%d ", r.Value)
	}
	// 4 3 2 1 0
	println()
	//使用Do调用匿名函数来打印环形链表的值
	r.Do(func(x interface{}) {
		fmt.Printf("%d ", x)
	})
	// 0 1 2 3 4
	println()
	//将环指针向后移动2步
	r = r.Move(2)
	r.Do(func(x interface{}) {
		fmt.Printf("%d ", x)
	})
	// 2 3 4 0 1
	println()
	//将环指针r后的第4个元素乘以5
	r.Move(4).Value = r.Move(4).Value.(int) * 5
	r.Do(func(x interface{}) {
		fmt.Printf("%d ", x)
	})
	// 2 3 4 0 5
	println()
	//删除环指针r后的3个元素
	res := r.Unlink(3)
	//可以看到原环形链表从r.Next()开始的3个元素被删除了
	r.Do(func(x interface{}) {
		fmt.Printf("%d ", x)
	})
	// 2 5
	println()
	//Unlink的返回值是被删除的元素组成的子环
	res.Do(func(x interface{}) {
		fmt.Printf("%d ", x)
	})
	// 3 4 0
	println()
	//将环形链表res插入到环形链表r的后面,返回值是插入前的r.Next()
	ans := r.Link(res)
	ans.Do(func(x interface{}) {
		fmt.Printf("%d ", x)
	})
	// 5 2 3 4 0
	println()
}
