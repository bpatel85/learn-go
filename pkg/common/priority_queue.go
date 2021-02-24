package common

// this is copied from https://golang.org/pkg/container/heap/

import (
	"container/heap"
)

// NewMinPriorityQueue returns the min priority queuq
func NewMinPriorityQueue() *PriorityQueue {
	pq := make(PriorityQueue, 0)
	return &pq
}

// An QueueItem is something we manage in a priority queue.
type QueueItem struct {
	Value    interface{} // The value of the item; arbitrary.
	Priority int         // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*QueueItem

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) IsEmpty() bool {
	return len(*pq) == 0
}

func (pq *PriorityQueue) HasNext() bool {
	return !pq.IsEmpty()
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*QueueItem)
	item.index = n
	*pq = append(*pq, item)

}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)

	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) Enqueue(x *QueueItem) {
	heap.Push(pq, x)
}

func (pq *PriorityQueue) Dequeue() (*QueueItem, error) {
	item := heap.Pop(pq).(*QueueItem)
	return item, nil
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) Update(item *QueueItem, value string, priority int) {
	item.Value = value
	item.Priority = priority
	heap.Fix(pq, item.index)
}
