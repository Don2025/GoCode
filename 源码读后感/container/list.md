## Container/list

Container中包含了Go常用的容器类型，本篇记录container/list的读后感。

### package list

list 是双端链表中的一个元素，空列表表示为nil列表，而零列表是一个具有零值的单元素列表。

```go
package list

// Element是链表中的一个元素
type Element struct {
	// next和prev是双端链表中指向前驱和后继元素的指针
	next, prev *Element
	// 此元素所属的列表
	list *List
	// 此元素所存储的值
	Value interface{}
}

// Next 返回下一个list元素或者nil
func (e *Element) Next() *Element {
	if p := e.next; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

// Prev 返回上一个list元素或者nil
func (e *Element) Prev() *Element {
	if p := e.prev; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

// List 表示双端链表
// List的零值是一个可以使用的空列表
type List struct {
	root Element // list的哨兵元素，仅使用&root, root.prev, root.next
	len  int     // 不包括哨兵元素的当前list的长度
}

// Init 用来初始化或者清空list.
/* 为了简化实现，可以在内部将list视为一个环，
 * &l.root既是最后一个元素l.Back()的后继元素
 * &l.root也是第一个元素l.Front()的前驱元素
**/
func (l *List) Init() *List {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0
	return l
}

// New 返回一个初始化的list
func New() *List { return new(List).Init() }

// Len 返回list中的元素个数，复杂度为 O(1)
func (l *List) Len() int { return l.len }

// Front 返回list的第一个元素，若list为空，则返回nil
func (l *List) Front() *Element {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}

// Back 返回list的最后一个元素，若list为空，则返回nil
func (l *List) Back() *Element {
	if l.len == 0 {
		return nil
	}
	return l.root.prev
}

// lazyInit 初始化一个零列表值
func (l *List) lazyInit() {
	if l.root.next == nil {
		l.Init()
	}
}

// insert 在at元素后插入e元素，并让l.len自增1，然后返回元素e
func (l *List) insert(e, at *Element) *Element {
	e.prev = at
	e.next = at.next
	e.prev.next = e
	e.next.prev = e
	e.list = l
	l.len++
	return e
}

// insertValue 是为了方便insert的一个包装器(&Element{Value: v}, at)
func (l *List) insertValue(v interface{}, at *Element) *Element {
	return l.insert(&Element{Value: v}, at)
}

// remove 从list中移除元素e，并让l.len自减，然后返回元素e
func (l *List) remove(e *Element) *Element {
	e.prev.next = e.next
	e.next.prev = e.prev
	e.next = nil // 避免内存泄漏
	e.prev = nil // 避免内存泄漏
	e.list = nil
	l.len--
	return e
}

// move 将元素e移动到元素at旁边并返回元素e
func (l *List) move(e, at *Element) *Element {
	if e == at {
		return e
	}
	e.prev.next = e.next
	e.next.prev = e.prev

	e.prev = at
	e.next = at.next
	e.prev.next = e
	e.next.prev = e

	return e
}

// Remove 能将list中的非nil元素e从list中移除，并返回元素值e.value
func (l *List) Remove(e *Element) interface{} {
	if e.list == l {
		// if e.list == l, l must have been initialized when e was inserted
		// in l or l == nil (e is a zero Element) and l.remove will crash
		l.remove(e)
	}
	return e.Value
}

// PushFront 能在list前面插入一个值为v的新元素e，并返回元素e
func (l *List) PushFront(v interface{}) *Element {
	l.lazyInit()
	return l.insertValue(v, &l.root)
}

// PushBack 能在list后面插入一个值为v的新元素e，并返回元素e
func (l *List) PushBack(v interface{}) *Element {
	l.lazyInit()
	return l.insertValue(v, l.root.prev)
}

// InsertBefore 能在非nil元素mark前插入一个值为v的新元素e，并返回元素e
// 若mark并不是list中的元素，则不修改list
func (l *List) InsertBefore(v interface{}, mark *Element) *Element {
	if mark.list != l {
		return nil
	}
	return l.insertValue(v, mark.prev)
}

// InsertAfter 能在非nil元素mark后插入一个值为v的新元素e，并返回元素e
// 如果mark不是list中的元素，则不修改list
func (l *List) InsertAfter(v interface{}, mark *Element) *Element {
	if mark.list != l {
		return nil
	}
	return l.insertValue(v, mark)
}

// MoveToFront 将非nil元素e移动到list的前面，若e不是list中的元素，则不修改list
func (l *List) MoveToFront(e *Element) {
	if e.list != l || l.root.next == e {
		return
	}
	l.move(e, &l.root)
}

// MoveToBack 将非nil元素e移动到list的后面，若e不是list中的元素，则不修改list
func (l *List) MoveToBack(e *Element) {
	if e.list != l || l.root.prev == e {
		return
	}
	l.move(e, l.root.prev)
}

// MoveBefore 将非nil元素e移动到非nil元素mark前的新位置
// 若e或mark不是list中的元素，或者e==mark则不修改list
func (l *List) MoveBefore(e, mark *Element) {
	if e.list != l || e == mark || mark.list != l {
		return
	}
	l.move(e, mark.prev)
}

// MoveAfter 将非nil元素e移动到非nil元素mark后的新位置
// 若e或mark不是list中的元素，或者e==mark则不修改list
func (l *List) MoveAfter(e, mark *Element) {
	if e.list != l || e == mark || mark.list != l {
		return
	}
	l.move(e, mark)
}

// PushBackList 在列表l后插入另一个列表other，这两个列表可以是一样的，但不能为nil
func (l *List) PushBackList(other *List) {
	l.lazyInit()
	for i, e := other.Len(), other.Front(); i > 0; i, e = i-1, e.Next() {
		l.insertValue(e.Value, l.root.prev)
	}
}

// PushFrontList 在列表l前插入另一个列表other，这两个列表可以是一样的，但不能为nil
func (l *List) PushFrontList(other *List) {
	l.lazyInit()
	for i, e := other.Len(), other.Back(); i > 0; i, e = i-1, e.Prev() {
		l.insertValue(e.Value, &l.root)
	}
}
```

### 用法示例

```go
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
```

### 例题求解

> #### 从上到下打印二叉树
>
> 从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印。
>
> 输入样例：[3, 9, 20, null, null, 15, 7]
>
> ```
>     3
>    / \
>   9  20
>     /  \
>    15   7
> ```
>
> 输出样例：[3, 9, 20, 15, 7]
>
> 来源：力扣（LeetCode）
>
> 链接：https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-lcof
>
> 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

比如树的层序遍历一般会用到队列来完成广度优先搜索，可以用Container/list来模拟queue的操作。

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) []int {
    var ans []int
    if root == nil {
        return ans
    }
    q := []*TreeNode {0: root}
    for len(q) != 0 {
        node := q[0]
        q = q[1:]
        ans = append(ans, node.Val)
        if node.Left != nil {
            q = append(q, node.Left)
        }
        if node.Right != nil {
            q = append(q, node.Right)
        }
    }
    return ans
}
```

因为Container中没有queue和stack，所以模拟栈和队列算是list的一个小用处吧，可以用Container/list来实现队列queue，

```go
package queue

import "container/list"

type Queue struct {
	list *list.List
}

// New returns new construct queue
func New() *Queue {
	return &Queue{list.New()}
}

// Push inserts element to the queue
func (queue *Queue) Push(value interface{}) {
	queue.list.PushBack(value)
}

// Front returns first element of the queue
func (queue *Queue) Front() interface{} {
	it := queue.list.Front()
	if it != nil {
		return it.Value
	}
	return nil
}

// Back returns last element of the queue
func (queue *Queue) Back() interface{} {
	it := queue.list.Back()
	if it != nil {
		return it.Value
	}
	return nil
}

// Pop returns and deletes first element of the queue
func (queue *Queue) Pop() interface{} {
	it := queue.list.Front()
	if it != nil {
		queue.list.Remove(it)
		return it.Value
	}
	return nil
}

// Size returns size of the queue
func (queue *Queue) Size() int {
	return queue.list.Len()
}

// Empty returns whether queue is empty
func (queue *Queue) Empty() bool {
	return queue.list.Len() == 0
}

// Clear clears the queue
func (queue *Queue) Clear() {
	for !queue.Empty() {
		queue.Pop()
	}
}
```

也可以用Container/list来实现栈stack。

```go
package stack

import "container/list"

type Stack struct {
    list *list.List
}

// New return new construct stack
func New *Stack {
    return &Stack{list.New()}
}

// Push insert element to the stack
func (stack *Stack) Push(value interface{}) {
    stack.list.PushBack(value)
}

// Pop return and delete top element of the stack
func (stack *Stack) Pop() interface{} {
    it := stack.list.Back()
    if it != nil {
        stack.list.Remove(it)
        return it.Value
    }
    return nil
}

// Top return top element of the stack
func (stack *Stack) Top() interface{} {
    it := stack.list.Back()
    if it != nil {
        return it.Value
    }
    return nil
}

// Size return size of the stack
func (stack *Stack) Size() int {
    return stack.list.Len()
}

// Empty return whether stack is empty
func (stack *Stack) Empty() bool {
    return stack.list.Len() == 0
}
```

