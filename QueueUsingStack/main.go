package main

import "fmt"

func main(){
	fmt.Println("Here we start!!")
	stack1:=stackStruct{stack:make([]int,5,5),top:-1,cap:5}
	stack2:=stackStruct{stack:make([]int,5,5),top:-1,cap:5}
	fmt.Println("Entering elements")
	for i:=1;i<=5;i++{
		stack1.push(i)
	}
	stack1.traverse()
	for i:=1;i<=5;i++{
		val,_:=stack1.pop()
		stack2.push(val)
	}
	stack2.pop()
	le:=stack2.top+1
	for i:=1;i<=le;i++{
		val,_:=stack2.pop()
		stack1.push(val)
	}
	fmt.Println("Printing stack 1")
	stack1.traverse()
}