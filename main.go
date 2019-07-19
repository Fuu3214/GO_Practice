package main

import "fmt"

func main() {
	intervals := [][]int{
		{1, 4}, {5, 8}, {9, 13}, {23, 77}, {78, 233},
	}
	newInterval := []int{4, 213}
	fmt.Println(insert(intervals, newInterval))
}
