package main

import "fmt"

/**
*	34. Find First and Last Position of Element in Sorted Array.
*	Given an array of integers nums sorted in non-decreasing order, find the starting and ending position of a given target value.
*	If target is not found in the array, return [-1, -1].
**/

/**
*	Runtime: 4 ms, faster than 98.27% of Go online submissions for Find First and Last Position of Element in Sorted Array.
*	Memory Usage: 4 MB, less than 58.20% of Go online submissions for Find First and Last Position of Element in Sorted Array.
**/

func main() {
	fmt.Println("Starting...........")
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	fmt.Printf("Array is %v and the target is %v and the result is : %v", nums, target, searchRange(nums, target))

}

func searchRange(nums []int, target int) []int {
	sindex := false
	findex := false
	indexes := []int{-1, -1}
	if len(nums) == 0 {
		return indexes
	}
	for i := 0; i <= len(nums)/2; i++ {
		if nums[i] > target {
			break
		}
		if nums[i] == target {
			indexes[0] = i
			sindex = true
		} else if nums[len(nums)-i-1] == target {
			indexes[1] = len(nums) - 1 - i
			findex = true
		}
		if sindex {
			for j := i; j < len(nums) && nums[j] == target; j++ {
				indexes[1] = j
			}
			break
		}
		if findex {
			for j := indexes[1]; j >= 0 && nums[j] == target; j-- {
				indexes[0] = j
			}
			break
		}
	}
	return indexes
}
