package common

import (
	"container/list"
	"errors"
	"sync"
)

type ThreadSafeQueue struct {
	lock sync.Mutex
	vals *list.List
}

func NewThreadSafeQueue() Queue {
	return ThreadSafeQueue{
		lock: sync.Mutex{},
		vals: list.New(),
	}
}

func (tsq ThreadSafeQueue) IsEmpty() bool {
	return tsq.vals.Len() == 0
}

func (tsq ThreadSafeQueue) HasNext() bool {
	return !tsq.IsEmpty()
}

func (tsq ThreadSafeQueue) Enqueue(val interface{}) {
	tsq.lock.Lock()
	defer tsq.lock.Unlock()

	tsq.vals.PushBack(val)
}

func (tsq ThreadSafeQueue) Dequeue() (interface{}, error) {
	tsq.lock.Lock()
	defer tsq.lock.Unlock()

	if tsq.vals.Len() == 0 {
		return nil, errors.New("Empty Stack")
	}

	front := tsq.vals.Front()
	tsq.vals.Remove(front)

	return front.Value, nil
}
