package solution9

type CQueue struct {
	Stack1 []int
	Stack2 []int
}

func Constructor() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	this.Stack1 = append(this.Stack1, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.Stack2) == 0 {
		if len(this.Stack1) == 0 {
			return -1
		}
		for len(this.Stack1) > 0 {
			v := this.Stack1[len(this.Stack1)-1]
			this.Stack2 = append(this.Stack2, v)
			this.Stack1 = this.Stack1[:len(this.Stack1)-1]
		}
	}

	res := this.Stack2[len(this.Stack2)-1]
	this.Stack2 = this.Stack2[:len(this.Stack2)-1]
	return res
}
