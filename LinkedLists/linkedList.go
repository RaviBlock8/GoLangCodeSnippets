package main

import (
	"errors"
	"fmt"
)

type linkedlist struct{
	head *node
}

type node struct{
	value string
	pointer *node
}

// Function below will take input from user and create a new node
func (ll *linkedlist) insert(value string){
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

func (ll *linkedlist) delete(value string) (*node,error){
	var deletedNode *node
	if (ll.head==nil){
		fmt.Println("List is empty , nothing to traverse!!!")
		return nil,errors.New("List is empty")
	}
	if(ll.head.value==value){
		deletedNode=ll.head
		ll.head=ll.head.pointer
		return deletedNode,nil
	}
	var previousNode *node
	var currentNode *node
	previousNode=ll.head
	currentNode=ll.head.pointer
	for currentNode!=nil{
		if currentNode.value==value{
			deletedNode=currentNode
			previousNode.pointer=currentNode.pointer
			return deletedNode,nil
		}
		previousNode=currentNode
		currentNode=currentNode.pointer
	}
	return nil,errors.New("Value not found in linked list")
}

func (ll *linkedlist) deleteFromEnd() (*node,error){
	var deletedNode *node
	if (ll.head==nil){
		fmt.Println("List is empty , nothing to traverse!!!")
		return nil,errors.New("List is empty")
	}
	var previousNode *node
	var currentNode *node
	previousNode=ll.head
	currentNode=ll.head.pointer
	for currentNode.pointer!=nil{
		previousNode=currentNode
		currentNode=currentNode.pointer
	}
	deletedNode=previousNode.pointer
	previousNode.pointer=nil
	return deletedNode,nil
}