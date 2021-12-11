package main

// Time Complexity: O(n^2)

func twoSum(nums []int, target int) []int {
	var out []int

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums) && j != i; j++ {
			if nums[i]+nums[j] == target {
				out = []int{i, j}
				break
			}
		}
	}

	return out
}
