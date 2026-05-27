type MinStack struct {
	stack []int
	min int
}

func Constructor() MinStack {
	return MinStack{
		min: 0,
	}
}

func (this *MinStack) Push(val int) {
	if len(this.stack)==0{
		this.min = val
		this.stack = append(this.stack, 0)
		return
	}
	this.stack = append(this.stack, val - this.min)
	if val < this.min { this.min = val}
}

func (this *MinStack) Pop() {
	pop := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	if pop < 0 {
		this.min = this.min - pop
	}
}

func (this *MinStack) Top() int {
	x := this.stack[len(this.stack)-1]
	if x<=0 {return this.min}
	return x+this.min
}

func (this *MinStack) GetMin() int {
	return this.min
}
