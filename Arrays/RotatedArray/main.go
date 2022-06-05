package main

import "fmt"

/**
*	33. Search in Rotated Sorted Array
*	Runtime: 5 ms, faster than 24.53% of Go online submissions for Search in Rotated Sorted Array.
*	Memory Usage: 2.9 MB, less than 16.78% of Go online submissions for Search in Rotated Sorted Array.
**/

func main() {
	fmt.Println("Starting................")
	nums := []int{1, 3}
	target := 3
	fmt.Printf("\nList of elements is : %v , target is : %v , and index is : %v", nums, target, search(nums, target))
}

func search(nums []int, target int) int {
	var goBackwards bool
	var goForwards bool

	// Here I am deciding , should I move forward or backward in terms of looking for position of element.
	if target >= nums[0] {
		goForwards = true
	} else if target <= nums[len(nums)-1] {
		goBackwards = true
	}
	// If it is neither greater then first element nor smaller then last element then it's not in the list return -1.
	if !goBackwards && !goForwards {
		return -1
	}

	// Now in case we want to move forward , traverse and look for element
	if goForwards {
		for i := 0; i < len(nums); i++ {
			if nums[i] < nums[0] || nums[i] > target {
				break
			}
			if nums[i] == target {
				return i
			}
		}
		return -1
	} else {
		for i := len(nums) - 1; i >= 0; i-- {
			if nums[i] < target || nums[i] > nums[len(nums)-1] {
				break
			}
			if nums[i] == target {
				return i
			}
		}
		return -1
	}
}
