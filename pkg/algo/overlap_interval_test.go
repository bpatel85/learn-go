package algo_test

import (
	"fmt"
	"testing"

	"github.com/bpatel85/learn-go/pkg/algo"
)

func TestOverlappingIntervals(t *testing.T) {
	input := []algo.Interval{
		{Start: 1, End: 2, Group: "a"},
		{Start: 3, End: 5, Group: "b"},
		{Start: 2, End: 3, Group: "a"},
		{Start: 6, End: 9, Group: "b"},
		{Start: 8, End: 9, Group: "c"},
		{Start: 7, End: 10, Group: "a"},
		{Start: 12, End: 13, Group: "b"},
	}

	a, err := algo.FindOverlappingIntervals(input, []string{"a", "b", "c"})
	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Printf("%+v", a)
}
