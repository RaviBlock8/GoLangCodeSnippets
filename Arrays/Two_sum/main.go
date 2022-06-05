package main

import "fmt"

/**
* Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
**/
/*
* Runtime: 14 ms, faster than 44.24% of Go online submissions for Two Sum.
* Memory Usage: 6.9 MB, less than 5.64% of Go online submissions for Two Sum.
 */

type hmap struct {
	index int
	pres  bool
}

func main() {
	fmt.Println("Starting...............")
	array := []int{2, 7, 11, 15}
	target := 9
	pair := twoSum(array, target)
	fmt.Printf("\nPair for this target : %v is  %v", target, pair)
}

func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]hmap)
	var pair []int
	for i, e := range nums {
		if !hashMap[e].pres {
			hashMap[e] = hmap{index: i, pres: true}
		}
	}
	for i, e := range nums {
		pending := target - e
		if hashMap[pending].pres {
			i2 := findIndex(i, pending, nums)
			if i2 == -1 {
				continue
			} else {
				return []int{i, i2}
			}
		}
	}
	return pair
}

func findIndex(i1 int, e2 int, nums []int) int {
	for i, e := range nums {
		if e == e2 && i != i1 {
			return i
		}
	}
	return -1
}
