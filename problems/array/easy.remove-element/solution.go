package main

// Time Complexity: O(n)

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	currentI := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[currentI] = nums[i]
			currentI++
		}
	}

	return currentI
}
