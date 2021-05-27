package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

func main() {
	var n, t, m int
	fmt.Scan(&n)
	head := &Node{}
	p := head
	for i := 0; i < n; i++ {
		fmt.Scan(&t)
		p.Next = &Node{t, nil}
		p = p.Next
	}
	fmt.Scan(&m)
	p = head
	for p.Next != nil {
		if p.Next.Data == m {
			p.Next = p.Next.Next
			n--
		} else {
			p = p.Next
		}
	}
	fmt.Println(n)
	p = head
	for p.Next != nil {
		fmt.Printf("%d ", p.Next.Data)
		p = p.Next
	}
}

/*
 * 运行时间：2ms 超过100.00%用Go提交的代码
 * 占用内存：920KB 超过60.00%用Go提交的代码
**/
