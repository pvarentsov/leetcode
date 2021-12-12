package main

import (
	"sort"
)

// Time Complexity: O(n^4*logn)
// Status: Passed

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)

	var res [][]int

	for i := 0; i < len(nums)-3; i++ {
		sumLeft := target - nums[i]
		res = treeSum(nums, res, sumLeft, i+1)
	}

	return res
}

func treeSum(nums []int, res [][]int, target int, startIndex int) [][]int {
	length := len(nums)

	for i := startIndex; i < len(nums)-2; i++ {
		sumLeft := target - nums[i]
		left := i + 1
		right := length - 1

		for left < right {
			sumRight := nums[left] + nums[right]

			if sumLeft == sumRight {
				quadruplet := []int{nums[startIndex-1], nums[i], nums[left], nums[right]}
				if !isDuplicate(res, quadruplet) {
					res = append(res, quadruplet)
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

func isDuplicate(res [][]int, quadruplet []int) bool {
	isDuplicate := false

	for i := len(res) - 1; i >= 0; i-- {
		value := res[i]

		var (
			first  = value[0] == quadruplet[0]
			second = value[1] == quadruplet[1]
			third  = value[2] == quadruplet[2]
			fourth = value[3] == quadruplet[3]
		)
		if first && second && third && fourth {
			isDuplicate = true
			break
		}
	}

	return isDuplicate
}
