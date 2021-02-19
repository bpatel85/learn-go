package algo

import (
	"errors"
	"math"
	"sort"
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

func FindOverlappingIntervals(intervals []Interval) ([]MergedInterval, error) {
	if len(intervals) == 0 {
		return nil, errors.New("empty intervals provided")
	}

	sortFunc := func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	}

	// todo: this is not recommended as this is mutating the passed in array
	sort.SliceStable(intervals, sortFunc)

	merged := make([]MergedInterval, 0)

	current := MergedInterval{
		Start:        intervals[0].Start,
		End:          intervals[0].End,
		SubIntervals: []Interval{intervals[0]},
	}

	for idx, i := range intervals {
		if idx == 0 {
			continue
		}

		if i.Start > current.End+1 {
			// this is disjoint
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

	merged = append(merged, current)

	return merged, nil
}
