package main

import "fmt"

type parenthesisString string

func (s parenthesisString)checkValidity() bool{
	st:=stackStruct{stack:make([]string,len(s),len(s)),top:-1,cap:len(s)}
	fmt.Println(st)
	for _,c:=range s{
		if st.top==-1{
			st.push(string(c))
		}else if ifMatches(st.stack[st.top],string(c))==true{
			st.pop()
		}else{
			st.push(string(c))
		}
	}
	
	return st.top==-1
}

func ifMatches(b1 string,b2 string)bool{
	var ans bool
	switch b1{
	case "(":
		if b2==")"{
			ans=true
		}
	case "[":
		if b2=="]"{
			ans=true
		}
	case "{":
		if b2=="}"{
			ans=true
		}
	}
	return ans
}