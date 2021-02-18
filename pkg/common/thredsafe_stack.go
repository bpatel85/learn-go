package common

import (
	"container/list"
	"errors"
	"sync"
)

type ThreadSafeStack struct {
	lock sync.Mutex
	vals *list.List
}

func NewThreadSaftStack() Stack {
	return ThreadSafeStack{
		lock: sync.Mutex{},
		vals: list.New(),
	}
}

func (tss ThreadSafeStack) Push(val interface{}) {
	tss.lock.Lock()
	defer tss.lock.Unlock()

	tss.vals.PushFront(val)
}

func (tss ThreadSafeStack) Peek() (interface{}, error) {
	return tss.topOfStack(false)
}

func (tss ThreadSafeStack) topOfStack(delAfterRead bool) (interface{}, error) {
	tss.lock.Lock()
	defer tss.lock.Unlock()

	length := tss.vals.Len()
	if length == 0 {
		return nil, errors.New("Empty Stack")
	}

	first := tss.vals.Front()
	if delAfterRead {
		tss.vals.Remove(first)
	}

	return first.Value, nil
}

func (tss ThreadSafeStack) Pop() (interface{}, error) {
	return tss.topOfStack(true)
}

func (tss ThreadSafeStack) IsEmpty() bool {
	return tss.vals.Len() == 0
}

func (tss ThreadSafeStack) HasNext() bool {
	return !tss.IsEmpty()
}
