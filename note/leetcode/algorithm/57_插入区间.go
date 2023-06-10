package algorithm

import (
	"sort"
)

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		intervals = append(intervals, newInterval)
		return intervals
	}

	leftPosition, rightPosition := 0, 0
	left := newInterval[0]
	right := newInterval[len(newInterval)-1]
	for i, item := range intervals {
		if item[0] > left {
			intervals = insertIntervals(intervals, []int{newInterval[0], 0}, i)
			leftPosition = i
			break
		} else if item[0] <= left && item[len(item)-1] >= left {
			leftPosition = i
			break
		} else if item[1] < left {
			if i == len(intervals)-1 {
				intervals = append(intervals, newInterval)
				leftPosition = i + 1
				break
			}

			continue
		}
	}

	for i := leftPosition; i < len(intervals); i++ {
		if right < intervals[i][0] {
			rightPosition = i - 1
			intervals[i-1][1] = right
			break
		} else if right == intervals[i][0] {
			rightPosition = i
			if intervals[i][1] == 0 {
				intervals[i][1] = right
			}
			break
		} else if right <= intervals[i][1] {
			rightPosition = i
			break
		} else {
			if i == len(intervals)-1 {
				intervals[i][1] = right
				rightPosition = len(intervals) - 1
			}
			continue
		}
	}
	if leftPosition != rightPosition {
		tmp := intervals[leftPosition][0]
		copy(intervals[leftPosition:], intervals[rightPosition:])
		intervals[leftPosition][0] = tmp
		intervals = intervals[:len(intervals)-(rightPosition-leftPosition)]
	}
	return intervals
}

func insertIntervals(intervals [][]int, newInterval []int, i int) [][]int {
	intervals = append(intervals, newInterval)
	copy(intervals[i+1:], intervals[i:])
	intervals[i] = newInterval
	return intervals
}

// {1,2}{4,6}{7,9}{10,12}  {5,8}
func insert1(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		intervals = append(intervals, newInterval)
		return intervals
	}

	if newInterval[0] < intervals[0][0] {
		return insertIntervals(intervals, newInterval, 0)
	}
	if newInterval[1] > intervals[len(intervals)-1][1] {
		return append(intervals, newInterval)
	}

	head := [][]int{}
	tail := [][]int{}
	start, end := 0, len(intervals)-1
	for {
		if newInterval[0] > intervals[start][1] {
			head = append(head, intervals[start])
			start++
			continue
		}
		if newInterval[1] < intervals[end][0] {
			tail = append(tail, intervals[end])
			end--
			continue
		}
		break
	}

	s := []int{newInterval[0], newInterval[1]}
	for i := start; i <= end; i++ {
		s = append(s, intervals[i]...)
	}
	sort.Ints(s)

	head = append(head, []int{s[0], s[len(s)-1]})
	reverseIntSlice(tail)
	head = append(head, tail...)
	return head
}

func reverseIntSlice(s [][]int) {
	for i := len(s)/2 - 1; i >= 0; i-- {
		opp := len(s) - 1 - i
		s[i], s[opp] = s[opp], s[i]
	}
}
