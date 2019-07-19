package main

func largestRectangleArea(heights []int) int {
	return largestRectangleArea2(heights)
}

func largestRectangleArea2(heights []int) int {
	return devideAndConqur(heights, 0, len(heights), 0)
}
func preprocess(heights []int, l int, r int) (int, bool) {
	min := l
	ordered := true
	for i := l + 1; i < r; i++ {
		if heights[min] > heights[i] {
			min = i
		}
		if heights[i-1] > heights[i] {
			ordered = false
		}
	}
	return min, ordered
}
func devideAndConqur(heights []int, l int, r int, area int) int {

	min, ordered := preprocess(heights, l, r)
	if ordered {
		for i := l; i < r; i++ {
			area = max(area, (r-i)*heights[i])

		}
		return area
	} else {
		area = max(area, heights[min]*(r-l))
		return max(devideAndConqur(heights, l, min, area), devideAndConqur(heights, min+1, r, area))
	}
}

func largestRectangleArea1(heights []int) int {
	s := make([]int, len(heights), len(heights))
	top := 0
	p := 0
	maxArea := 0
	var w int

	for p < len(heights) {
		if top == 0 || heights[s[top-1]] <= heights[p] {
			s[top] = p
			top++
			p++
		} else {
			pop := s[top-1]
			top--
			if top == 0 {
				w = p
			} else {
				w = p - s[top-1] - 1
			}
			maxArea = max(maxArea, w*heights[pop])
		}
	}

	//dump the stack
	for top > 0 {
		pop := s[top-1]
		top--
		if top == 0 {
			w = p
		} else {
			w = p - s[top-1] - 1
		}
		maxArea = max(maxArea, w*heights[pop])
	}

	return maxArea
}
