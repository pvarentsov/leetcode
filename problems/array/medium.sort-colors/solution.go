package main

// Time Complexity: O(n^2)
// Algorithm: Shaker Sort
// Status: Passed

func sortColors(nums []int) {
	left := 0
	right := len(nums) - 1

	for left <= right {
		for i := left; i < right; i++ {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
		}

		right -= 1

		for i := right; i > left; i-- {
			if nums[i-1] > nums[i] {
				nums[i], nums[i-1] = nums[i-1], nums[i]
			}
		}

		left += 1
	}
}
