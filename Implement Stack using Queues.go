//225. Implement Stack using Queues
//https://leetcode.com/problems/implement-stack-using-queues/
type MyStack struct {
	Stack []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{Stack: []int{}}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.Stack = append(this.Stack, x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	pop := this.Stack[len(this.Stack)-1]
	this.Stack = this.Stack[:len(this.Stack)-1]
	return pop
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.Stack[len(this.Stack)-1]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	if len(this.Stack) == 0 {
		return true
	}
	return false
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */