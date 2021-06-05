## Container/ring

Container中包含了Go常用的容器类型，本篇记录`container/ring`的读后感。

### package ring

ring是环形链表中的一个元素，环没有起点和终点，指向任何ring元素的指针都是对整个环的引用。空环表示为nil环指针，而零环是一个具有零值的单元素环。

```go
package ring

type Ring struct {
  next, prev *Ring
  Value interface{}
}

// init 初始化一个单元素环
func (r *Ring) init() *Ring {
	r.next = r
	r.prev = r
	return r
}

// Next 返回非空环元素r的下一个环元素
func (r *Ring) Next() *Ring {
	if r.next == nil {
		return r.init()
	}
	return r.next
}

// Prev 返回非空环元素r的前一个环元素
func (r *Ring) Prev() *Ring {
	if r.next == nil {
		return r.init()
	}
	return r.prev
}

// Move 能将非空环r中的n%r.Len()个元素向后移动(n<0),或向前移动(n>=0),并返回该环元素
func (r *Ring) Move(n int) *Ring {
	if r.next == nil {
		return r.init()
	}
	switch {
	case n < 0:
		for ; n < 0; n++ {
			r = r.prev
		}
	case n > 0:
		for ; n > 0; n-- {
			r = r.next
		}
	}
	return r
}

// New 创建一个包含n个元素的环
func New(n int) *Ring {
	if n <= 0 {
		return nil
	}
	r := new(Ring)
	p := r
	for i := 1; i < n; i++ {
		p.next = &Ring{prev: p}
		p = p.next
	}
	p.next = r
	r.prev = p
	return r
}

// Link 能将非空环r与环元素s连接起来，使r.Next()变成环s，并返回r.Next()的原始值
/* 如果环r和环s指向同一个环，将它们连接起来就会把环元素r和环元素s之间的元素从环中移除，
 * 被删除的元素形成一个子环，返回结果是对该子环的引用。
 * 如果没有被删除的元素，返回结果仍是r.Next()的原始值，而不是nil。
**/
/* 如果环r和环s指向不同的环，则在环元素r后连接环元素s从而创建一个单独的环，
 * 返回结果指向环元素s插入后的最后一个元素的下一个元素。
**/
func (r *Ring) Link(s *Ring) *Ring {
	n := r.Next()
	if s != nil {
		p := s.Prev()
    //不能在一行多重赋值，这样会打乱执行顺序
		r.next = s
		s.prev = r
		n.prev = p
		p.next = n
	}
	return n
}

/* Unlink 能从非空环r中移除从r.next()开始的n%r.Len()个环元素，若n%r.Len()==0,环r不变
 * 返回值是被移除元素构成的子环。
**/
func (r *Ring) Unlink(n int) *Ring {
	if n <= 0 {
		return nil
	}
	return r.Link(r.Move(n + 1))
}

// Len 返回环r中的元素个数，执行时间与元素数成正比
func (r *Ring) Len() int {
	n := 0
	if r != nil {
		n = 1
		for p := r.Next(); p != r; p = p.next {
			n++
		}
	}
	return n
}

// Do 对环r中的每个元素依次调用函数f，如果f改变环指针*r，则Do的行为是未定义的
func (r *Ring) Do(f func(interface{})) {
	if r != nil {
		f(r.Value)
		for p := r.Next(); p != r; p = p.next {
			f(p.Value)
		}
	}
}
```

