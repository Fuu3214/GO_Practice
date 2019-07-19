package main

import "math"

func max(A int, B int) int {
	return int(math.Max(float64(A), float64(B)))
}
