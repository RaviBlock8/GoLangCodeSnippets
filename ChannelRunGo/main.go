package main

import (
	"fmt"
	"time"
)

func main(){
	// /**
	// * So as we know in go whenever you create a variable using the keyword var.
	// * It gets initialized by 0 equivalent of it.
	// * That's why incase of reference type , it is initialized by nil.
	// * Same with channels just like in slices , if you create one using var.
	// * It will create nil channel as shown below.
	// **/
	// var c chan int
	// <-c
	/**
	* So in can check status of channel as shown below.
	* val,ok:=<-c
	* So in go lang it provides a syntax using which it will move out of for loop , as soon as channel is closed , it will come out of for loop.
	* for val:=range c{fmt.Println(val)}
	**/
	/**
	* --------- BUFFERED CHANNEL---------------
	* So in normal channel , when you push a value , channel is blocked until that value is received on the other side.
	* In buffer channel , a buffer of particular size is created , so when you push values in channel , it is pushed into buffer.
	* And when you try to read values from channel on another side then these values are read at one go.
	* So this is created using this make(chan int,10).
	* So in buffered channel , it will only block when you will send values more then the capacity of the channel, until the values are read.
	**/
	// //------------UNIDIRECTIONAL CHANNEL-------------
	// c:=make(chan int)
	// go readData(c)
	// for i:=0;i<6;i++{
	// 	c<-i
	// }
	// close(c)
	// soc:=make(chan int)
	// go sendData(soc)
	// for val:=range soc{
	// 	fmt.Println("SOC Value:",val)
	// }

	// //------------USING CHANNEL AS AN FIRST HAND DATA TYPE---------
	// cc:=make(chan chan int)
	// go doubleChannel(cc)
	// c:=<-cc
	// for val:=range c{
	// 	fmt.Println("Value read in double channel:",val)
	// }
	/**
	* ------SELECT STATEMENT----------
	* So select statement is equivalent to switch case but for channels.
	* So this select statement is used when you are dealing with multiple channels.
	* So in this , it will stay a blocking operation , unless you have put a default case.
	* Incase of default case , it will keep executing default case , until some channel becomes unblocking.
	* In case if all the channels are unblocking then it will randomly pick one channel and get the value and then move on to the next one.
	**/
		// c1:=blockingOne()
		// c2:=blockingTwo()
		// for{
		// 	switch{
		// 	case val1 := <-c1:
		// 		fmt.Println(val1)
		// 	case val2 := <-c2:
		// 		fmt.Println(val2)
		// 	}
		// }

}

func readData(c <-chan int){
	for val:=range c{
		fmt.Println("Value read:",val)
	}
}

func sendData(c chan<-int){
	for i:=0;i<6;i++{
		c<-i
	}
	close(c)
}

func doubleChannel(cc chan chan int){
	c:=make(chan int)
	cc<-c
	for i:=0;i<6;i++{
		c<-i
	}
	close(c)

}

func blockingOne() <-chan int{
	c:=make(chan int)
	for i:=0;i<6;i++{
		c<-i
		time.Sleep(time.Duration(3)*time.Second)
	}
	return c
}

func blockingTwo() <-chan int{
	c:=make(chan int)
	for i:=0;i<6;i++{
		time.Sleep(time.Duration(3)*time.Second)
		c<-i
	}
	return c
}