package common_test

import (
	"testing"

	"github.com/bpatel85/learn-go/pkg/common"
)

func TestPriorityQueue(t *testing.T) {
	q := common.NewPriorityQueue()

	if !q.IsEmpty() {
		t.Errorf("newly created pq must be empty")
	}

	if q.HasNext() {
		t.Errorf("newly created pq must not have any vars")
	}

	for i := 0; i < 5; i++ {
		itm := common.QueueItem{
			Value:    i,
			Priority: i,
		}
		q.Enqueue(&itm)
	}

	for i := 0; i < 5; i++ {
		ret, err := q.Dequeue()
		if err != nil {
			t.Errorf("error in dequeueing %v", err)
		}

		if ret.Priority != i {
			t.Errorf("pq dequeue mismatch. Expected %v, got %v", i, ret.Priority)
		}
	}

	if !q.IsEmpty() {
		t.Errorf("pq must be empty after all dq")
	}
}
