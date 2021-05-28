//155. Min Stack
//https://leetcode.com/problems/min-stack/
type MinStack struct {
	Stack []int
	Min   int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{Stack: make([]int, 0)}
}

func (this *MinStack) Push(val int) {
	this.Stack = append(this.Stack, val)
	if len(this.Stack) == 1 {
		this.Min = val
	}
	if this.Min > val {
		this.Min = val
	}
}

func (this *MinStack) Pop() {
	if this.Stack[len(this.Stack)-1] == this.Min {
		this.Min = this.Stack[0]
		for i := 1; i < len(this.Stack)-1; i++ {
			if this.Stack[i] < this.Min {
				this.Min = this.Stack[i]
			}
		}
	}
	this.Stack = this.Stack[:len(this.Stack)-1]
}

func (this *MinStack) Top() int {
	return this.Stack[len(this.Stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.Min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */