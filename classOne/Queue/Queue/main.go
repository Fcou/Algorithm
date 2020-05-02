package main

//根据status设计队列,链表结构,前出后进，使用前替换status
type myQueue struct {
	first *myQueueNode
	end   *myQueueNode
}

type myQueueNode struct {
	val  status
	next *myQueueNode
}

func (queue *myQueue) init() {
	queue.first = nil
	queue.end = nil
}

func (queue *myQueue) Push(n status) {
	if n == nil {
		return
	}
	newNode := &myQueueNode{
		val:  n,
		next: nil,
	}
	if queue.first == nil && queue.end == nil {
		queue.first = newNode
		queue.end = newNode
		return
	}
	queue.end.next = newNode
	queue.end = newNode
}

func (queue *myQueue) Pop() (q status) {
	if queue.first == queue.end {
		q = queue.first.val
		queue.first = nil
		queue.end = nil
		return
	}
	q = queue.first.val
	queue.first = queue.first.next
	return
}

func (queue *myQueue) empty() bool {
	if queue.end == nil {
		return true
	}
	return false
}
