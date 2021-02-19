package algo_test

import (
	"fmt"
	"testing"

	"github.com/bpatel85/learn-go/pkg/algo"
)

func TestOverlappingIntervals(t *testing.T) {
	input := []algo.Interval{
		algo.Interval{Start: 1, End: 2, Group: "a"},
		algo.Interval{Start: 3, End: 5, Group: "b"},
		algo.Interval{Start: 2, End: 3, Group: "a"},
		algo.Interval{Start: 6, End: 9, Group: "b"},
		algo.Interval{Start: 8, End: 9, Group: "c"},
		algo.Interval{Start: 7, End: 10, Group: "a"},
		algo.Interval{Start: 12, End: 13, Group: "b"},
	}

	a, err := algo.FindOverlappingIntervals(input)
	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Printf("%+v", a)
}
