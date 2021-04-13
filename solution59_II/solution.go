package solution59_II

type MaxQueue struct {
	data []int
	max  []int
}

func Constructor() MaxQueue {
	return MaxQueue{[]int{}, []int{}}
}

func (this *MaxQueue) Max_value() int {
	if len(this.max) == 0 {
		return -1
	}
	return this.max[0]
}

func (this *MaxQueue) Push_back(value int) {

	for len(this.max) > 0 && this.max[len(this.max)-1] < value {
		this.max = this.max[:len(this.max)-1]
	}
	this.max = append(this.max, value)
	this.data = append(this.data, value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.data) == 0 {
		return -1
	}
	front := this.data[0]
	if len(this.max) > 0 && this.max[0] == front {
		this.max = this.max[1:]
	}
	this.data = this.data[1:]
	return front
}
