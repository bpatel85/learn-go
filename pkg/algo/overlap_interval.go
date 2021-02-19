package algo

import (
	"errors"
	"math"
	"sort"

	"github.com/bpatel85/learn-go/pkg/common"
)

/** a single interval is defined as a 3-element tuple
# (lower_bound, upper_bound, category)

# task:
# implement a method find_overlapping_intervals that takes two inputs
#     input 1: a list of tuples representing intervals (described above)
#     input 2: a tuple of categories
# and returns
#     a list where each element is a list of mutually overlapping intervals that belong to any of the input categories (see example below)
**/

type Interval struct {
	Start, End int
	Group      string
}

type MergedInterval struct {
	Start, End   int
	SubIntervals []Interval
}

func FindOverlappingIntervals(intervals []Interval, categories []string) ([]MergedInterval, error) {
	if len(categories) == 0 {
		return nil, nil
	}

	if len(intervals) == 0 {
		return nil, errors.New("empty intervals provided")
	}

	filtered := make([]Interval, 0)
	for _, i := range intervals {
		if common.StringInSliceIgnoreCase(i.Group, categories) {
			filtered = append(filtered, i)
		}
	}

	if len(filtered) == 0 {
		return nil, nil
	}

	sortFunc := func(i, j int) bool {
		return filtered[i].Start < filtered[j].Start
	}

	// sort and initial state
	sort.SliceStable(filtered, sortFunc)
	merged := make([]MergedInterval, 0)
	current := MergedInterval{
		Start:        filtered[0].Start,
		End:          filtered[0].End,
		SubIntervals: []Interval{filtered[0]},
	}

	for idx, i := range filtered {
		if idx == 0 {
			continue
		}

		if i.Start > current.End+1 {
			// this is disjoint. start new merged interval
			merged = append(merged, current)
			current = MergedInterval{
				Start:        i.Start,
				End:          i.End,
				SubIntervals: []Interval{i},
			}
		} else {
			current.End = int(math.Max(float64(i.End), float64(current.End)))
			current.SubIntervals = append(current.SubIntervals, i)
		}
	}

	// add the last stage
	merged = append(merged, current)

	return merged, nil
}
