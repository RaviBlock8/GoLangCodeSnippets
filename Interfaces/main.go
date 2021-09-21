package main

import "fmt"

/**
* So in this code below we will first create a code with Bad patterns and then we will see how we can resolve then using interfaces.
* Here below we have two struct and we have created a receiver function which will be called in reference to these struct.
* One interesting point to note here is that even though function name is same for both the cases.
* Go lang considers them different just because reference type is different.
**/

type ebBot struct{}

type spanishBot struct{}

// So this here it means , that " Hey all the types in this program there is a new type by name bot".
// And whoever type in this program will have method getGreeting which will return a string.
// Will automatically be promoted to honorary type bot.
// And will be able to use printGreeting.
type bot interface{
	getGreeting() string
}

func main(){
	e:=ebBot{}
	s:=spanishBot{}
	printGreeting(e)
	printGreeting(s)
}
func (ebBot) getGreeting() string{
	return "Hi there!!"
}

func (spanishBot) getGreeting() string{
	return "Hola!!!!"
}

/**
* Unlike reference functions , here go lang is giving error because name is same , even though args type is different.
**/

func printGreeting(b bot){
	fmt.Println(b.getGreeting())
}

/**
* So unlike other languages we don't have to manually tell go that.
* Hey ebBot implements bot interface.
* Go identifies that by itself.
* Following are some points regarding interfaces.
* 1. Interfaces are not generic types in go.
* 2. Interfaces are 'implicit'.
* 3. Interfaces are a contract to help us manage types.
* 4. Interfaces are tough. Step#1 is understanding how to read them.
**/