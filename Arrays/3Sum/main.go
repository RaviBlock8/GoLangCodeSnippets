package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
)

// Problem with this approach is that it is more of brute force approach and takes lots of time.

func main() {
	fmt.Println("Starting")
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Printf("\n List is : %v , permutations are : %v", nums, threeSum(nums))
}

func threeSum(nums []int) [][]int {
	var perm [][]int
	hashToValue := make(map[string][]int)
	hashToFound := make(map[string]bool)

	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums); j++ {
			for k := 2; k < len(nums); k++ {
				if i != j && i != k && j != k && (nums[i]+nums[j]+nums[k]) == 0 {
					comb := []int{nums[i], nums[j], nums[k]}
					cpyComb := []int{comb[0], comb[1], comb[2]}
					bubbleSort(comb)
					hash := getHash(comb)
					if !hashToFound[hash] {
						hashToValue[hash] = cpyComb
						hashToFound[hash] = true
					}
				}
			}
		}
	}

	for key, val := range hashToValue {
		fmt.Printf("\n Hash : %v \n", key)
		fmt.Printf("\n Value : %v \n", val)
		perm = append(perm, val)
	}
	return perm
}

func bubbleSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		for j := 0; j < len(nums)-i; j++ {
			if nums[j] > nums[j+1] {
				tem := nums[j+1]
				nums[j+1] = nums[j]
				nums[j] = tem
			}
		}
	}
}

func getHash(nums []int) string {
	byteSlice, err := json.Marshal(nums)
	if err != nil {
		panic(("error while marshalling"))
	}
	encodedSlice := base64.StdEncoding.EncodeToString(byteSlice)
	return encodedSlice
}
