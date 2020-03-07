//container/heap 中优先级队列的源代码
// This example demonstrates a priority queue built using the heap interface.
package main

import (
	"container/heap"
	"fmt"
)

// hnode 哈夫曼树节点定义
type hnode struct {
	node   string  // 编码
	weight float32 // 权重
	number int     // 编号,原始节点编号从0到n-1，新生产节点从n开始，总共会有n-1个
	left   *hnode  //左孩子
	right  *hnode  //右孩子
}

// PriorityQueue 优先级队列，实现了接口，内含元素为hnode的指针
type PriorityQueue []*hnode

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].weight > pq[j].weight
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].number = i
	pq[j].number = j
}

//Push 向队列中加入新节点
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*hnode)
	item.number = n
	*pq = append(*pq, item)
}

//Pop 弹出最小节点
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}

// 更新节点
func (pq *PriorityQueue) update(item *hnode, node string, weight float32) {
	item.node = node
	item.weight = weight
	heap.Fix(pq, item.number)
}

//根据*hnode设计队列,链表结构,前出后进
type myQueue struct {
	first *myQueueNode
	end   *myQueueNode
}

type myQueueNode struct {
	val  *hnode
	next *myQueueNode
}

func (queue *myQueue) init() {
	queue.first = nil
	queue.end = nil
}

func (queue *myQueue) Push(n *hnode) {
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
	return
}

func (queue *myQueue) Pop() (q *hnode) {
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

func main() {
	//P 为26个英文字母出现的统计概率（来源维基百科）,因为要利用最小堆，要转换一下统计概率
	P := []float32{0.08167, 0.01492, 0.02782, 0.04253, 0.12702, 0.02228, 0.02015, 0.06094, 0.06966, 0.00153, 0.00772, 0.04025, 0.02406, 0.06749, 0.07507, 0.01929, 0.00095, 0.05987, 0.06327, 0.09056, 0.02758, 0.00978, 0.02360, 0.00150, 0.01974, 0.00074}
	for i, v := range P {
		P[i] = 1 - v
	}
	//创建一个节点队列，用P初始化，构建26个字母根据概率排序的最小优先级队列
	var size int
	if len(P) > 0 {
		size = 2*len(P) - 1 //最终会用到n+n-1个节点
	} else {
		return
	}
	D := make([]hnode, size)
	index := 0
	pq := make(PriorityQueue, 0) //注意刚开始队列中没元素，要设为0，以后通过Push添加
	heap.Init(&pq)
	for key, probability := range P {
		D[key] = hnode{
			node:   "",
			weight: probability,
			number: key,
			left:   nil,
			right:  nil,
		}
		heap.Push(&pq, &D[key])
		index++
	}

	//利用已有的优先级队列，创建新节点，构建哈夫曼树，最终优先级队列中剩余根节点
	for pq.Len() > 1 {
		item := heap.Pop(&pq).(*hnode)
		D[index].number = index
		D[index].weight = item.weight
		D[index].left = item
		item2 := heap.Pop(&pq).(*hnode)
		D[index].weight += item2.weight
		D[index].right = item2
		heap.Push(&pq, &D[index])
		index++
	}
	//利用队列，层次遍历之前构建的哈夫曼树,写入对应节点的编码
	var Q myQueue
	Q.init()
	if pq.Len() != 0 {
		root := heap.Pop(&pq).(*hnode)
		Q.Push(root)
	}
	for !Q.empty() {
		current := Q.first
		if current.val.left == nil {
			break //如果左孩子为空，说明已到最下层，停止
		}
		current.val.left.node = current.val.node + "0"
		current.val.right.node = current.val.node + "1"
		Q.Pop()
		Q.Push(current.val.left)
	}
	//输出A-Z对应的哈夫曼编码
	for i := 0; i < len(P); i++ {
		fmt.Printf("%c 的哈夫曼编码是：%s\n", byte('A'+i), D[i].node)
	}
}
