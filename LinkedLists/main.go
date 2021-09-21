package main

import "fmt"

func main(){
	fmt.Println("Linked list program\n")
	var ll linkedlist
	ll.insert("1")
	ll.insert("2")
	ll.insert("3")
	ll.insert("4")
	ll.insert("5")
	ll.insert("6")
	ll.insert("7")
	ll.traverse()
	fmt.Println("Deleting node with value 4")
	ll.delete("4")
	ll.traverse()
	fmt.Println("Deleting node from end")
	ll.deleteFromEnd()
	ll.traverse()
}