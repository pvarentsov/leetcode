package main

import (
	"sort"
)

// Time Complexity: O(n^3)
// Status: Time Limit Exceeded

func threeSum(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	if len(nums) == 1 && nums[0] == 0 {
		return [][]int{}
	}

	var out [][]int
	target := 0

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums) && j != i; j++ {
			for k := 0; k < len(nums) && k != i && k != j; k++ {

				if nums[i]+nums[j]+nums[k] == target {
					triplet := []int{nums[i], nums[j], nums[k]}
					sort.Ints(triplet)

					if len(out) == 0 {
						out = append(out, triplet)
						continue
					}

					isDuplicate := false

					for _, value := range out {
						if value[0] == triplet[0] && value[1] == triplet[1] && value[2] == triplet[2] {
							isDuplicate = true
							break
						}
					}

					if !isDuplicate {
						out = append(out, triplet)
					}
				}
			}
		}
	}

	return out
}
