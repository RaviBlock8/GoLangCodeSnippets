package main

import "fmt"

func main(){
	var sampleString parenthesisString
	sampleString="[]("
	if sampleString.checkValidity(){
		fmt.Println("Yes it has valid parenthesis")
	}else{
		fmt.Println("No it doesn't have valid parenthesis")
	}
}