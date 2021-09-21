package main

import (
	"fmt"
	"net/http"
)

/**
* So by default we have only one go routine and that is main go routine.
* To do executions in parallel what you can do is create multiple go routines in parallel.
* For that we use go keyword.
* So whenever you call a function to with go keyword , a separate go routine is spin up.
* And assigned to that function.
**/

/**
* ------Concurrency and parallelism--------
* So go by default uses one CPU core to run go routines.
* Even incase of multiple routines , it switches back and forth b/w them.This is handled by go scheduler.
* So when one routine is waiting for some i/o , go scheduler will pick another one.
* This is basically what concurrency is.
* Parallelism is when we run these go routines in multiple cores at same time.
**/

/**
* -----------Channels-----------
* So go lang by default creates main go routine.
* Main routine decides how long program will persists.
* So in code below when we use go keyword to create go routines.
* Nothing happens , because main thread comes out of for loop and finishes
* So one way to communicate between go routines is channel.
* Using channels we will establish connection between main go routine and other child routines.
**/

/**
* -----------Using for loop with Channels----------
* for val:=range c{} so it is waiting for channel c to return value and then for loop iterate
**/
func main(){
	fmt.Println("Start")
	links:=[]string{"http://google.com","http://facebook.com","http://stackoverflow.com"}
	c:=make(chan string)
	no:=0

	for _,link:=range(links){
		go checkLink(link,c)
	}
	for true{
		fmt.Println(<-c)
		no+=1
		if no==len(links){
			break
		}
	}
}

// So we are using this function to check these links sequentially
// To do it parallely we can use go routines. 
func checkLink(link string,c chan string){
	_,err:=http.Get(link)
	if err!=nil{
		fmt.Println(link,"Might be down")
		c<-"Stop"
		return
	}
	fmt.Println(link,"is up")
	c<-"Stop"
}