package main

import "testing"

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

/*
*
Input
["MyQueue", "push", "push", "peek", "pop", "empty"]
[[], [1], [2], [], [], []]
Output
[null, null, null, 1, 1, false]

Explanation
MyQueue myQueue = new MyQueue();
myQueue.push(1); // queue is: [1]
myQueue.push(2); // queue is: [1, 2] (leftmost is front of the queue)
myQueue.peek(); // return 1
myQueue.pop(); // return 1, queue is [2]
myQueue.empty(); // return false
*
*/
func TestMyQueue(t *testing.T) {
	myQueue := Constructor()
	myQueue.Push(1)
	myQueue.Push(2)
	peek := myQueue.Peek()
	if peek != 1 {
		t.Errorf("got %v, wanted %v", peek, 1)
	}
	pop := myQueue.Pop()
	if pop != 1 {
		t.Errorf("got %v, wanted %v", pop, 1)
	}
	empty := myQueue.Empty()
	if empty {
		t.Errorf("got %v, wanted %v", empty, false)
	}
}
