// This example demonstrates a priority queue built using the heap interface.
// 根据官方代码改成，dist最小堆
package dijkstra

import (
	"container/heap"
)

// An Item is something we manage in a priority queue.
type Item struct {
	id   int // The id of the item; arbitrary.
	dist int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].dist < pq[j].dist
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
func (pq *PriorityQueue) update(item *Item, id int, dist int) {
	item.id = id
	item.dist = dist
	heap.Fix(pq, item.index)
}

// func TestExample_priorityQueue(t *testing.T) {
// 	// Some items and their priorities.
// 	items := map[int]int{
// 		0: 3, 1: 11, 2: 4,
// 	}
// 	// Create a priority queue, put the items in it, and
// 	// establish the priority queue (heap) invariants.
// 	pq := make(PriorityQueue, len(items))
// 	i := 0
// 	for id, dist := range items {
// 		pq[i] = &Item{
// 			id:    id,
// 			dist:  dist,
// 			index: i,
// 		}
// 		i++
// 	}
// 	heap.Init(&pq)
// 	// // Insert a new item and then modify its priority.
// 	// item := &Item{
// 	// 	id:   3,
// 	// 	dist: 1,
// 	// }
// 	// heap.Push(&pq, item)
// 	pq.update(pq[0], 0, 99)
// 	// Take the items out; they arrive in decreasing priority order.
// 	for pq.Len() > 0 {
// 		item := heap.Pop(&pq).(*Item)
// 		fmt.Printf("dist:%d-id:%d    ", item.dist, item.id)
// 	}
// 	// Output:
// 	// 05:orange 04:pear 03:banana 02:apple
// }
