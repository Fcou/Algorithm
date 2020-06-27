package circularqueue

//循环队列
type CircularQueue struct {
	items []string
	n     int //队列容量
	head  int
	tail  int
}

func (this *CircularQueue) newCircularQueue(n int) *CircularQueue {
	this.items = make([]string, n)
	this.n = n
	this.head = 0
	this.tail = 0
	return this
}

func (this *CircularQueue) push(item string) bool {
	if (this.tail+1)%this.n == this.head {
		return false
	}
	this.items[this.tail] = item
	this.tail = (this.tail + 1) % this.n
	return true
}

func (this *CircularQueue) pop() (result string, ok bool) {
	if this.head == this.tail {
		result = ""
		ok = false
		return
	}
	result = this.items[this.head]
	ok = true
	this.head = (this.head + 1) % this.n
	return
}
