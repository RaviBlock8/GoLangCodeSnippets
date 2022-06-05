package main

import "fmt"

/**
	* 35. Search Insert Position
	* Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return
	* the index where it would be if it were inserted in order.
**/

func main() {
	fmt.Println("Starting.........")
	nums := []int{1, 2, 3}
	target := 1
	fmt.Printf("Answer:%v", searchInsert(nums, target))
}

func searchInsert(nums []int, target int) int {
	indexWhereItCouldBe := -1
	for i := 0; i < len(nums) && nums[i] <= target; i++ {
		if nums[i] == target {
			return i
		}
		indexWhereItCouldBe = i
	}
	return indexWhereItCouldBe + 1
}
