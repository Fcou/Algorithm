// This example demonstrates a priority queue built using the heap interface.
// 根据官方代码改成，f最小堆
package dijkstra

import (
	"container/heap"
)

// An Item is something we manage in a priority queue.
type Item struct {
	id int // The id of the item; arbitrary.
	f  int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	// 我们需要最小优先队列
	return pq[i].f < pq[j].f
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and id of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, id int, f int) {
	item.id = id
	item.f = f
	heap.Fix(pq, item.index)
}
