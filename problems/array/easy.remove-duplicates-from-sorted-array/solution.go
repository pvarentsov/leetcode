package main

// Time Complexity: O(n)

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	currentI := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[currentI] {
			currentI++
			nums[currentI] = nums[i]
		}
	}

	return currentI + 1
}
