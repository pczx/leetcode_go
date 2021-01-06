package solution30

import "math"

/** initialize your data structure here. */
type MinStack struct {
	dataStack []int
	minStack  []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{[]int{}, []int{math.MaxInt64}}
}

func (this *MinStack) Push(x int) {
	this.dataStack = append(this.dataStack, x)
	if len(this.minStack) == 0 || x <= this.minStack[len(this.minStack)-1] {
		this.minStack = append(this.minStack, x)
	}
}

func (this *MinStack) Pop() {
	top := this.dataStack[len(this.dataStack)-1]
	this.dataStack = this.dataStack[:len(this.dataStack)-1]
	if top == this.minStack[len(this.minStack)-1] {
		this.minStack = this.minStack[:len(this.minStack)-1]
	}
}

func (this *MinStack) Top() int {
	return this.dataStack[len(this.dataStack)-1]

}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
