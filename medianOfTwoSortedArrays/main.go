package main

import "fmt"

func main(){
	fmt.Println("Starting program")
	nums1:=[]int{1,2,3,4,6,7}
	nums2:=[]int{2,8,9}
	findMedianSortedArrays(nums1,nums2)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) {
    len1:=len(nums1)
	len2:=len(nums2)
	nums3:=make([]int,len1+len2,len1+len2)
	i1:=0
	i2:=0
	i3:=0
	for{
		if i1==len1{
			break
		}
		if i2==len2{
			break
		}
		if nums1[i1]==nums2[i2]{
			nums3[i3]=nums1[i1]
			i3+=1
			nums3[i3]=nums2[i2]
			i1+=1
			i2+=1
		}else if nums1[i1]<nums2[i2]{
			nums3[i3]=nums1[i1]
			i1+=1
		}else{
			nums3[i3]=nums2[i2]
			i2+=1
		}
		i3+=1
	}
	if i1!=len1{
		for i:=i1;i<=len1-1;i++{
			nums3[i3]=nums1[i]
			i3+=1
		}
	}
	if i2!=len2{
		for i:=i2;i<=len2-1;i++{
			nums3[i3]=nums2[i]
			i3+=1
		}
	}
	fmt.Println(nums3)
}