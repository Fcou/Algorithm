//container/heap 中优先级队列的源代码
// This example demonstrates a priority queue built using the heap interface.
package main

import (
	"container/heap"
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

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*hnode)
	item.number = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil   // avoid memory leak
	item.number = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *hnode, node string, weight float32) {
	item.node = node
	item.weight = weight
	heap.Fix(pq, item.number)
}

func main() {
	//P 为26个英文字母出现的统计概率（来源维基百科）
	P := []float32{0.08167, 0.01492, 0.02782, 0.04253, 0.12702, 0.02228, 0.02015, 0.06094, 0.06966, 0.00153, 0.00772, 0.04025, 0.02406, 0.06749, 0.07507, 0.01929, 0.00095, 0.05987, 0.06327, 0.09056, 0.02758, 0.00978, 0.02360, 0.00150, 0.01974, 0.00074}

	//创建一个节点队列，用P初始化，构建26个字母根据概率排序的最小优先级队列
	var size int
	if len(P) > 0 {
		size = 2*len(P) - 1 //最终会用到n+n-1个节点
	} else {
		return
	}
	D := make([]hnode, size)
	index := 0
	pq := make(PriorityQueue, len(P))
	for key, probability := range P {
		D[key] = hnode{
			node:   "",
			weight: probability,
			number: key,
			left:   nil,
			right:  nil,
		}
		heap.Push(&D[key])
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
	//最终优先级队列中剩余根节点

}
