package main

func insert(intervals [][]int, newInterval []int) [][]int {
	idxLL := find(intervals, newInterval[0], 0)
	idxLR := find(intervals, newInterval[0], 1)
	idxRL := find(intervals, newInterval[1], 0)
	idxRR := find(intervals, newInterval[1], 1)

	var newL int
	if idxLL == idxLR {
		// newInterval[0] should be between 2 intervals, after idxLL or idxLR
		if idxLL >= 0 && intervals[idxLL][1] == newInterval[0] {
			//should merge with idxLL
			intervals[idxLL][1] = newInterval[1]
			newL = idxLL
		} else if idxLL < len(intervals)-1 && newInterval[1] >= intervals[idxLL+1][0] {
			//shoudl merge with idxLL + 1
			intervals[idxLL+1][0] = newInterval[0]
			intervals[idxLL+1][1] = max(intervals[idxLL+1][1], newInterval[1])
			newL = idxLL + 1
		} else {
			// shoudl insert a new interval
			return sliceInsert(intervals, idxLL+1, newInterval)
		}
	} else {
		// newInterval[0] should be inside 1 intervals, in idxLL
		intervals[idxLL][1] = max(intervals[idxLL][1], newInterval[1])
		newL = idxLL
	}

	if idxRL == idxRR {
		// newInterval[1] should be between 2 intervals, after idxRL or idxRR
		return sliceRemove(intervals, newL+1, idxRL)
	} else {
		// newInterval[1] should be inside 1 intervals, inside idxRL
		intervals[newL][1] = intervals[idxRL][1]
		return sliceRemove(intervals, newL+1, idxRL)
	}
}

func sliceInsert(slice [][]int, index int, value []int) [][]int {
	if index == len(slice) {
		return append(slice, value)
	}
	slice = append(slice[:index+1], slice[index:]...)
	slice[index] = value
	return slice
}

func sliceRemove(slice [][]int, from int, to int) [][]int {
	if from <= to {
		slice = append(slice[:from], slice[to+1:]...)
	}
	return slice
}

func find(intervals [][]int, tar int, LR int) int {
	l := 0
	r := len(intervals) - 1
	idx := -1
	for l <= r {
		mid := (r-l)/2 + l
		if intervals[mid][LR] <= tar {
			idx = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return idx
}
