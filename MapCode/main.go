package main

import "fmt"

/**
* --------MAP--------------
* In this , we will learn about maps in go lang.
* Maps in go lang are like any other maps , which is basically a key value pair.
* There are multiple ways to declare it in go lang as shown below.
* var colors map[string]string
* colors:=map[string]string{"red":"#1234","blue":"#4523"}
* colors:=make(map[string]string)
**/

func main(){
	colors:=map[string]string{
		"red":"#123456",
		"blue":"#345667",
	}
	fmt.Println(colors)
	for key,value:=range(colors){
		fmt.Println(key)
		fmt.Println(value)
	}
	colors["red"]="78765"
	fmt.Println(colors)
}