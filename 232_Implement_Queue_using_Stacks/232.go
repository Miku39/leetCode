package main

import "container/list"

// 232. Implement Queue using Stacks
// https://leetcode.com/problems/implement-queue-using-stacks/

// TODO: use stack
type MyQueue struct {
	v *list.List
}

func Constructor() MyQueue {
	return MyQueue{
		v: list.New(),
	}
}

func (this *MyQueue) Push(x int) {
	this.v.PushBack(x)
}

func (this *MyQueue) Pop() int {
	if this.Empty() {
		return -1
	}
	first := this.v.Front()
	this.v.Remove(first)
	return first.Value.(int)
}

func (this *MyQueue) Peek() int {
	if this.Empty() {
		return -1
	}
	return this.v.Front().Value.(int)
}

func (this *MyQueue) Empty() bool {
	if this.v.Len() == 0 {
		return true
	}
	return false
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
