package main

import (
	"errors"
	"fmt"
)

type stackStruct struct{
	stack []int
	top int
	cap int
}

func (s *stackStruct) push(val int)(bool,error){
	if(s.top==(s.cap-1)){
		return false,errors.New("Overflow occured")
	}
	s.top+=1;
	s.stack[s.top]=val
	return true,nil
}

func (s *stackStruct) pop()(int,error){
	if(s.top==-1){
		return -1,errors.New("Underflow occured")
	}
	popped:=s.stack[s.top]
	s.top-=1;
	return popped,nil
}

func (s *stackStruct) traverse(){
	if(s.top==-1){
		fmt.Println("Stack is empty")
		return
	}
	for i:=s.top;i>-1;i-=1{
		fmt.Println("Value in stack:",s.stack[i])
	}
}