//232. Implement Queue using Stacks
//https://leetcode.com/problems/implement-queue-using-stacks/
type MyQueue struct {
	Queue []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{Queue: []int{}}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.Queue = append(this.Queue, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if len(this.Queue) == 0 {
		return 0
	}
	pop := this.Queue[0]
	this.Queue = this.Queue[1:]
	return pop
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	return this.Queue[0]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	if len(this.Queue) == 0 {
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