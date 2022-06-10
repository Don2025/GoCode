package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	e1 := l.PushFront(1)
	e4 := l.PushBack(4)
	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)
	fmt.Printf("%d\n", l.Len())
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%d ", e.Value)
	}
	println()
	for e := l.Front(); e != nil; e = e.Next() {
		if e.Value == 3 {
			e3 := l.Remove(e)
			fmt.Printf("remove: %d\n", e3.(int))
		}
	}
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%d ", e.Value)
	}
	println()
}
