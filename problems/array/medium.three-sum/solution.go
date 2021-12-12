package main

import (
	"sort"
)

// Time Complexity: O(n^3*logn)
// Status: Passed

func threeSum(nums []int) [][]int {
	target := 0
	length := len(nums)

	if length == 0 {
		return [][]int{}
	}
	if length == 1 && nums[0] == 0 {
		return [][]int{}
	}

	var res [][]int
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		sumLeft := target - nums[i]
		left := i + 1
		right := length - 1

		for left < right {
			sumRight := nums[left] + nums[right]

			if sumLeft == sumRight {
				triplet := []int{nums[i], nums[left], nums[right]}

				if !isDuplicate(res, triplet) {
					res = append(res, triplet)
				}

				left++
				right--
			}
			if sumLeft > sumRight {
				left++
			}
			if sumLeft < sumRight {
				right--
			}
		}
	}

	return res
}

func isDuplicate(res [][]int, triplet []int) bool {
	isDuplicate := false

	for i := len(res) - 1; i >= 0; i-- {
		value := res[i]

		if value[0] == triplet[0] && value[1] == triplet[1] && value[2] == triplet[2] {
			isDuplicate = true
			break
		}
	}

	return isDuplicate
}
