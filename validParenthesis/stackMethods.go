package main

import (
	"errors"
	"fmt"
)

type stackStruct struct{
	stack []string
	top int
	cap int
}

func (s *stackStruct) push(val string)(bool,error){
	if(s.top==(s.cap-1)){
		fmt.Println("Overflow occured!")
		return false,errors.New("Overflow")
	}
	s.top+=1;
	s.stack[s.top]=val
	return true,nil
}

func (s *stackStruct)pop()(string,error){
	if(s.top==-1){
		fmt.Println("Underflow occured!")
		return "",errors.New("Underflow")
	}
	popped:=s.stack[s.top]
	s.top-=1
	return popped,nil
}

func (s *stackStruct)traverse(){
	if(s.top==-1){
		fmt.Println("Stack is empty")
		return
	}
	for i:=s.top;i>=0;i--{
		fmt.Println("Value at stack:",s.stack[i])
	}
}