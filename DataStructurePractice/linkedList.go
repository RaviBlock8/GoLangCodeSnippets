package main

import "fmt"

type node struct {
	value int
	next  *node
}

type linkedList struct {
	head *node
}

func (ll *linkedList) InsertAtEnd(val int) bool {
	newNode := node{value: val, next: nil}
	if ll.head == nil {
		ll.head = &newNode
		return true
	}
	pointer := ll.head
	for pointer.next != nil {
		pointer = pointer.next
	}
	pointer.next = &newNode
	return true
}

func (ll *linkedList) InsertAtBegin(val int) bool {
	newNode := node{value: val, next: nil}
	if ll.head == nil {
		ll.head = &newNode
		return true
	}
	newNode.next = ll.head
	ll.head = &newNode
	return true
}

func (ll *linkedList) deleteBegin() int {
	if ll.head == nil {
		fmt.Errorf("Underflow condition")
		return -1
	}
	deletedNode := ll.head
	ll.head = ll.head.next
	return deletedNode.value
}

func (ll *linkedList) deleteEnd() int {
	if ll.head == nil {
		fmt.Errorf("Underflow condition")
		return -1
	}
	if ll.head.next == nil {
		ll.head = nil
	}
	pointer := ll.head
	for pointer.next.next != nil {
		pointer = pointer.next
	}
	deletedNode := pointer.next
	pointer.next = nil
	return deletedNode.value
}
