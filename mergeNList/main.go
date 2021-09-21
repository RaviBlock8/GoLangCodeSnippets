package main

import "fmt"

func main(){
	fmt.Println("Start program")
	listsOfLists:=[][]int{[]int{1,4,5},[]int{1,3,4},[]int{1,4,5}}
	fmt.Println("Lists of list:",listsOfLists)
	var ll linkedlist
	for index,list:=range listsOfLists{
		if index==0{
			for _,val:=range list{
				fmt.Println("Inserting:",val)
				ll.insert(val)
			}
		}else{
			fmt.Println("Merging list:",list)
			ll.mergeAnotherList(list)
		}
	}
	fmt.Println("-----------SORTED LIST NOW---------------------")
	ll.traverse()
}