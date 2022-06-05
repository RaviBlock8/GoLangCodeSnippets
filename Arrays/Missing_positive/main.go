package main

import "fmt"

//https://leetcode.com/problems/first-missing-positive/submissions/

func main() {
	nums := []int{-1, -2, -60, 40, 43}
	fmt.Println(nums)
	fmt.Println(bubbleSort(nums))
	fmt.Printf("\nMissing:%v", firstMissingPositive(bubbleSort((nums))))
}

func bubbleSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[j] > nums[j+1] {
				temp := nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = temp
			}
		}
	}
	return nums
}

func firstMissingPositive(sorted []int) int {
	missing := 0
	smallest := sorted[0]
	i := 0
	for ; smallest != sorted[len(sorted)-1]; smallest++ {
		if smallest != sorted[i] && smallest > 0 {
			if sorted[0] > 1 {
				missing = 1
			} else {
				missing = smallest
			}
			break
		}
		if smallest != sorted[i] && smallest <= 0 {
			continue
		}
		i++
	}
	if missing == 0 {
		if sorted[0] == 1 || sorted[0] == 0 {
			missing = sorted[len(sorted)-1] + 1
		} else {
			missing = 1
		}
	}
	return missing
}
