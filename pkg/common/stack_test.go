package common_test

import (
	"testing"

	"github.com/bpatel85/learn-go/pkg/common"
)

func TestStack(t *testing.T) {
	stk := common.NewThreadSaftStack()

	if !stk.IsEmpty() {
		t.Errorf("newly created stack must be empty")
	}

	if stk.HasNext() {
		t.Errorf("newly created stack must not have any vars")
	}

	for i := 0; i < 5; i++ {
		stk.Push(i)
	}

	for i := 4; i >= 0; i-- {
		ret, err := stk.Pop()
		if err != nil {
			t.Errorf("error in poping %v", err)
		}

		if ret.(int) != i {
			t.Errorf("pop mismatch. Expected %v, got %v", i, ret)
		}
	}

	if !stk.IsEmpty() {
		t.Errorf("statck must be empty after all pop")
	}
}
