package main

import "fmt"

type tnode struct {
	data   interface{}
	left   *tnode
	right  *tnode
	parent *tnode
}

//树的前中后三种递归遍历方法-属于深度优先遍历
//pre_order 前序遍历，先根，然后左子树，最后右子树
func pre_order(treeRootNode *tnode) {
	if treeRootNode != nil {
		fmt.Println(treeRootNode.data)
		pre_order(treeRootNode.left)
		pre_order(treeRootNode.right)
	}
}

//in_order 前序遍历，先左子树，然后根，最后右子树
func in_order(treeRootNode *tnode) {
	if treeRootNode != nil {
		in_order(treeRootNode.left)
		fmt.Println(treeRootNode.data)
		in_order(treeRootNode.right)
	}
}

//post_order 前序遍历，先左子树，然后右子树，最后根
func post_order(treeRootNode *tnode) {
	if treeRootNode != nil {
		post_order(treeRootNode.left)
		post_order(treeRootNode.right)
		fmt.Println(treeRootNode.data)
	}
}

//层次遍历方法，利用队列，广度优先搜索算法
func level_order(treeRootNode *tnode) {
	var Q myQueue
	if treeRootNode != nil {
		Q.Push(treeRootNode)
	} else {
		return
	}
	for !Q.empty() {
		f := Q.first
		fmt.Println(f.val.data)
		if f.val.left != nil {
			Q.Push(f.val.left)
		}
		if f.val.right != nil {
			Q.Push(f.val.right)
		}
		Q.Pop()
	}
}

//根据tnode设计队列,链表结构,前出后进
type myQueue struct {
	first *myQueueNode
	end   *myQueueNode
}

type myQueueNode struct {
	val  *tnode
	next *myQueueNode
}

func (queue *myQueue) init() {
	queue.first = nil
	queue.end = nil
}

func (queue *myQueue) Push(n *tnode) {
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

func (queue *myQueue) Pop() (q *tnode) {
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
