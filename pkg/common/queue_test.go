package common_test

import (
	"testing"

	"github.com/bpatel85/learn-go/pkg/common"
)

func TestQueue(t *testing.T) {
	q := common.NewThreadSafeQueue()

	if !q.IsEmpty() {
		t.Errorf("newly created q must be empty")
	}

	if q.HasNext() {
		t.Errorf("newly created q must not have any vars")
	}

	for i := 0; i < 5; i++ {
		q.Enqueue(i)
	}

	for i := 0; i < 5; i++ {
		ret, err := q.Dequeue()
		if err != nil {
			t.Errorf("error in dequeueing %v", err)
		}

		if ret.(int) != i {
			t.Errorf("dequeue mismatch. Expected %v, got %v", i, ret)
		}
	}

	if !q.IsEmpty() {
		t.Errorf("q must be empty after all dq")
	}
}
