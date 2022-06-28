package _41

// https://leetcode.cn/problems/flatten-nested-list-iterator/
//------------------------Leetcode Problem 341------------------------
/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */

type NestedIterator struct {
	// 将列表视作一个队列，栈中直接存储该队列
	stack [][]*NestedInteger
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	return &NestedIterator{[][]*NestedInteger{nestedList}}
}

func (it *NestedIterator) Next() int {
	// 由于保证调用 Next 之前会调用 HasNext，直接返回栈顶列表的队首元素，将其弹出队首并返回
	queue := it.stack[len(it.stack)-1]
	val := queue[0].GetInteger()
	it.stack[len(it.stack)-1] = queue[1:]
	return val
}

func (it *NestedIterator) HasNext() bool {
	for len(it.stack) > 0 {
		queue := it.stack[len(it.stack)-1]
		if len(queue) == 0 { // 当前队列为空，出栈
			it.stack = it.stack[:len(it.stack)-1]
			continue
		}
		nest := queue[0]
		if nest.IsInteger() {
			return true
		}
		// 若队首元素为列表，则将其弹出队列并入栈
		it.stack[len(it.stack)-1] = queue[1:]
		it.stack = append(it.stack, nest.GetList())
	}
	return false
}

//------------------------Leetcode Problem 341------------------------
/*
 * https://leetcode.cn/problems/flatten-nested-list-iterator/
 * 执行用时：4ms 在所有Go提交中击败了85.64%的用户
 * 占用内存：5.1MB 在所有Go提交中击败了100.00%的用户
**/
