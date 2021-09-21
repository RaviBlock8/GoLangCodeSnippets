package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
)

/**
* So here we creating anew data type by name deck.
* This data type is basically an slice of strings.
* The reason we are creating a new type is , because then we can add more functionalities to this type.
**/

type deck []string

/**
* So now here we are extending the functionality of this deck type.
* By creating this receiver function.
* It will receive variable of deck type.
* And apply operations on it
* Just like methods in javascript.
* So now any variable of type deck have access to this method print.
* So as per the convention in the receiver name of  variable should be first few letters of data type.
**/
func (d deck) print() {
	// This is how we iterate over a slice
	for index,cardElement:=range(d){
		fmt.Println(index,cardElement)
	}
}

func newDeck() deck{
	cardSuits:=deck{"Spades","Hearts","Diamonds"}
	cardValues:=deck{"Ace","Two","Three","Four"}
	var cardDeck deck
	for _,suit:=range(cardSuits){
		for _,value:=range(cardValues){
			cardDeck=append(cardDeck,value+" of "+suit)
		}
	}
	return cardDeck;
}

func deal(d deck,handSize int) (deck,deck){
	return d[:handSize],d[handSize:]
}

func (d deck) toString() string{
	// So strings is one of the go packages to deal with operations related to string
	return strings.Join([]string(d),",")
}

func (d deck) writeDataToFile(filename string) error{
	// ioutil is go package that deals with input output operations.
	//Like reading and writing to file.
	return ioutil.WriteFile(filename,[]byte(d.toString()),0666)
}

func newDeckFromFile(filename string) (deck,error){
	// So here I am reading data from file and I will get data in bytes format
	bts,err:=ioutil.ReadFile(filename)
	// err will not be nil if something wrong happened.
	if err!=nil{
		return nil,err
	}
	// Here we will first convert bytes into string.
	// Then we will split that string using , as separator.
	// Then I will convert it into deck type
	return deck(strings.Split(string(bts),",")),nil
}

func (d deck) shuffle() (deck,error){
	for i,_:=range(d){
		randIndex:= rand.Intn(len(d)-1)
		rep:=d[randIndex]
		d[randIndex]=d[i]
		d[i]=rep
	}
	return d,nil
}