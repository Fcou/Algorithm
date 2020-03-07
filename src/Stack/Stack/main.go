package main

//根据status设计栈,链表结构,使用前替换status
type myStack struct {
	first *myStackNode
	end   *myStackNode
}

type myStackNode struct {
	val  status
	next *myStackNode
}

func (stack *myStack) Push(n status) {
	newNode := &myStackNode{
		val:  n,
		next: nil,
	}
	if stack.first == nil && stack.end == nil {
		stack.first = newNode
		stack.end = newNode
		return
	}
	stack.end.next = newNode
	stack.end = newNode
}

func (stack *myStack) Pop() (q status) {

	if stack.empty() { //空栈返回false
		return
	}
	q = stack.end.val

	tem := stack.first
	if stack.first == stack.end {
		stack.first = nil
		stack.end = nil
	}
	for {
		if tem.next == stack.end {
			stack.end = tem
			break
		}
		tem = tem.next
	}
	return
}

func (stack *myStack) empty() bool {
	if stack.first == nil {
		return true
	}
	return false
}
