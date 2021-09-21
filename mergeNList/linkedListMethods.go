package main

import (
	"fmt"
)

type linkedlist struct{
	head *node
}

type node struct{
	value int
	pointer *node
}

// Function below will take input from user and create a new node
func (ll *linkedlist) insert(value int){
	newNode:=node{value:value,pointer: nil}
	var customNode *node
	if(ll.head==nil){
		ll.head=&newNode
		fmt.Printf("New node created:%+v\n",newNode)
		return
	}else{
		customNode=ll.head
		for customNode.pointer!=nil{
			customNode=customNode.pointer
		}
		fmt.Printf("New node created:%+v\n",newNode)
		customNode.pointer=&newNode
	}
}

func (ll *linkedlist) traverse(){
	if (ll.head==nil){
		fmt.Println("List is empty , nothing to traverse!!!")
	}else{
		fmt.Println("Traversing linked list")
		var customNode *node
		customNode=ll.head
		for customNode!=nil{
			fmt.Printf("Node value:%+v\n",customNode.value)
			customNode=customNode.pointer
		}
		fmt.Println("Traveresing is done!!!")
	}
}

func (ll *linkedlist)mergeAnotherList(list []int){
	for _,val:=range list{
		ll.insertAtSortedPosition(val)
	}
}

func (ll *linkedlist) insertAtSortedPosition(val int){
	pointer:=ll.head
	var previous *node
	newNode:=node{value: val,pointer: nil}
	for pointer!=nil{
		if(pointer.value>val){
			if previous==nil{
				ll.head=&newNode
				newNode.pointer=pointer
			}else{
				previous.pointer=&newNode
				newNode.pointer=pointer
			}
			fmt.Println("Element entered:",val);
			return
		}
		previous=pointer
		pointer=pointer.pointer
	}
	previous.pointer=&newNode
}