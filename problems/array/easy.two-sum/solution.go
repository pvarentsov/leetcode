package main

// Time Complexity: O(n^2)
// Status: Passed

func twoSum(nums []int, target int) []int {
	var res []int

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums) && j != i; j++ {
			if nums[i]+nums[j] == target {
				res = []int{i, j}
				break
			}
		}
	}

	return res
}
