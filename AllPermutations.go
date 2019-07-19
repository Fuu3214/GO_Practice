package main

func permute(nums []int) [][]int {
	ret := make([][]int, 0)
	permuteImpl(nums, len(nums), &ret)
	return ret
}

func permuteImpl(nums []int, n int, ret *[][]int) {
	if n == 1 {
		row := make([]int, len(nums), len(nums))
		copy(row, nums)
		*ret = append(*ret, row)
	} else {
		for i := 0; i < n; i++ {
			swap(nums, i, n-1)
			permuteImpl(nums, n-1, ret)
			swap(nums, i, n-1)
		}
	}
}

func permuteImplOpt(nums []int, n int, ret *[][]int) {
	if n == 1 {
		row := make([]int, len(nums), len(nums))
		copy(row, nums)
		*ret = append(*ret, row)
	} else {
		for i := 0; i < n-1; i++ {
			permuteImplOpt(nums, n-1, ret)
			if n%2 == 0 {
				swap(nums, i, n-1)
			} else {
				swap(nums, 0, n-1)
			}

		}
		permuteImplOpt(nums, n-1, ret)
	}
}

func swap(nums []int, a int, b int) {
	nums[a], nums[b] = nums[b], nums[a]
}
