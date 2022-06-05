package main

import "fmt"

/**
* Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.
* Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
* Output: 6
* Explanation: [4,-1,2,1] has the largest sum = 6.
**/

/**
* ------------- DISCUSSION ----------------
* The easiest approach we can think of is brute force , where I try to form every contiguous subarray and store with maximum output.
* Problem is that will take 3 loops , loop 1 -> subarray size , loop 2 -> Goes through each element , loop 3 -> create subarray starts from jth element of size i.
* Hint says divide and conqeor can be used so think about that.
**/

func main() {
	nums := []int{-1}
	fmt.Printf("Maximum value of contiguous sub array sum : %v", maxSubArray(nums))
}

func maxSubArray(nums []int) int {
	max := 0
	for i := range nums {
		curSum := maxPossibleValue(nums, i)
		if curSum > max {
			max = curSum
		}
	}
	return max
}

func maxPossibleValue(nums []int, ci int) int {
	max := nums[ci]
	currentSum := 0
	for i := ci; i < len(nums); i++ {
		currentSum += nums[i]
		if currentSum > max {
			max = currentSum
		}
	}
	return max
}
