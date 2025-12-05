package day5

import "sort"

type IdRange struct {
	Left  int
	Right int
}

func MergeIntervals(intervals []IdRange) []IdRange {
	if len(intervals) == 0 {
		return nil
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Left < intervals[j].Left
	})

	merged := []IdRange{intervals[0]}

	for _, interval := range intervals[1:] {
		prevInterval := &merged[len(merged)-1]

		// Overlapping intervals
		if interval.Left <= prevInterval.Right+1 {
			if interval.Right > prevInterval.Right {
				prevInterval.Right = interval.Right
			}
		} else {
			merged = append(merged, interval)
		}
	}

	return merged
}

func InRangeSearch(id int, intervals []IdRange) bool {
	i := sort.Search(len(intervals), func(i int) bool {
		return intervals[i].Left > id
	})
	i--
	if i < 0 {
		return false
	}
	return intervals[i].Left <= id && id <= intervals[i].Right
}
