package merge_intervals

import "sort"

type Interval struct {
	Start int
	End   int
}

type interSlice []Interval

func (s interSlice) Len() int { return len(s) }

func (s interSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s interSlice) Less(i, j int) bool { return s[i].Start < s[j].Start }

func merge(intervals []Interval) []Interval {
	if len(intervals) <= 1 {
		return intervals
	}
	sort.Sort(interSlice(intervals))

	result := make([]Interval, 0, len(intervals))
	start := intervals[0].Start
	end := intervals[0].End

	for i := 1; i < len(intervals); i++ {
		inte := intervals[i]
		if inte.Start <= end { // 合并
			if inte.End > end {
				end = inte.End
			}
		} else { // 不能合并，把前面的当做一个间隔
			result = append(result, Interval{start, end})
			start, end = inte.Start, inte.End
		}
	}
	result = append(result, Interval{start, end})
	return result
}
