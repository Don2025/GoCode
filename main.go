package main

//test github.com/

import (
	stack "github.com/Don2025/golib/stack"
	"testing"
)

func main() {
	var t *testing.T = &testing.T{}
	testStack(t)
}

func testStack(t *testing.T) {
	s := stack.NewStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	size := s.Size()
	if size != 4 {
		t.Errorf("stack.Size() test failed! Got %d, expected 4.\n", size)
	}
	value := s.Top().(int)
	if value != 4 {
		t.Errorf("stack.Top() test failed! Got %d, expected 4.\n", value)
	}
	value = s.Pop().(int)
	if value != 4 {
		t.Errorf("stack.Pop() test failed! Got %d, expected 4.\n", value)
	}
	size = s.Size()
	if size != 3 {
		t.Errorf("stack.Size() test failed! Got %d, expected 3.\n", size)
	}
	value = s.Pop().(int)
	if value != 3 {
		t.Errorf("stack.Pop() test failed! Got %d, expected 3.\n", value)
	}
	value = s.Top().(int)
	if value != 2 {
		t.Errorf("stack.Top() test failed! Got %d, expected 2.\n", value)
	}
	empty := s.Empty()
	if empty {
		t.Errorf("stack.Empty() test failed! Got %t, expected false.\n", empty)
	}
	value = s.Pop().(int)
	if value != 2 {
		t.Errorf("stack.Pop() test failed! Got %d, expected 2.\n", value)
	}
	empty = s.Empty()
	if empty {
		t.Errorf("stack.Empty() test failed! Got %t, expected false.\n", empty)
	}
	value = s.Pop().(int)
	if value != 1 {
		t.Errorf("stack.Pop() test failed! Got %d, expected 1.\n", value)
	}
	empty = s.Empty()
	if !empty {
		t.Errorf("stack.Empty() test failed! Got %t, expected true.\n", empty)
	}
	nilValue := s.Top()
	if nilValue != nil {
		t.Errorf("stack.Top() test failed! Got %d, expected nil.", nilValue)
	}
	nilValue = s.Pop()
	if nilValue != nil {
		t.Errorf("stack.Pop() test failed! Got %d, expected nil.", nilValue)
	}
	size = s.Size()
	if size != 0 {
		t.Errorf("stack.Size() test failed! Got %d, expected 0.", size)
	}
}
