package algo_test

import (
	"testing"

	"github.com/bpatel85/learn-go/pkg/algo"
)

func TestZK(t *testing.T) {
	zk := algo.NewZK()
	if err := zk.Create("/broker1/topics/sensorA", "hello"); err != nil {
		t.Fail()
	}

	if data, err := zk.GetData("/broker1/topics/sensorA"); err != nil || data != "hello" {
		t.Fail()
	}

	if _, err := zk.GetData("/broker1/topics/sensorB"); err == nil {
		t.Fail()
	}

	if err := zk.SetData("/broker1/topics/sensorA", "foo"); err != nil {
		t.Fail()
	}

	if data, err := zk.GetData("/broker1/topics/sensorA"); err != nil || data != "foo" {
		t.Fail()
	}

	if depth := zk.Depth(); depth != 3 {
		t.Fail()
	}

	if err := zk.SetData("/", "foo"); err != nil {
		t.Fail()
	}

	if exists := zk.Exists("/"); !exists {
		t.Fail()
	}

	if data, err := zk.GetData("/"); err != nil || data != "foo" {
		t.Fail()
	}

	if deleted := zk.Delete("/foobar"); deleted {
		t.Fail()
	}

	if deleted := zk.Delete("/broker1/topics"); !deleted {
		t.Fail()
	}

	if err := zk.SetData("/broker1/topics/sensorA", "foo"); err == nil {
		t.Fail()
	}
}
