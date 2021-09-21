package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
* -----Go routines-----
* It is independently executing function launched by go.
* It's very cheap , it's very practical to have thousands of go routines.
* There might be only one thread in a program with thousands of routines.
* But if you think of it as a very cheap thread , it won't be  far.
**/

/**
* ----Channels-----
* So channels in go are used for the communication between 2 go routines.
* So these channels are first hand data type in go lang.
* So in channel when a routine uses channel to receive a value it is an blocking operation.
* But also when it sends the value it is an blocking operation , and it waits for receiver to receive it.
* So channels not only helps to communicate but it also allows for synchronization.
* Buffer channels are also there in go lang that doesn't synchronize and just throws the value in buffer.
* Go approach is " Don't communicate by sharing memory , share memory by communicating.".
**/

func main(){
	// //----------GENERATOR PATTERN--------
	// c:=generatorPattern("Generator")
	// for{
	// 	fmt.Println(<-c)
	// }

	// //-------------CHANNELS AS SERVICE------------
	// joe:=generatorPattern("Joe")
	// alice:=generatorPattern2("Alice")
	// /**
	// * So here in channel as service pattern we are creating multiple go routines.
	// * So in this pattern these are taking turns not only in printing the value but also in executing.
	// * Because of sync nature of channels , if alice is ready to return.
	// * But joe isn't done yet , alice will never be able to print value.
	// * To counter this we have multiplexed pattern or fan in.
	// **/
	// // Now the interesting part again is that alice is channel which have more messages to send.
	// // But go scheduler running it round robin fashion.
	// // So let's say the first channel is taking some time in it's turn. Main thread is not letting alice speak.
	// // So alice have to wait even if he have stuff to say and joe is just thinking.
	// for{
	// 	fmt.Println(<-joe)
	// 	fmt.Println(<-alice)
	// }

	// //--------FAN IN PATTERN--------
	// /**
	// * So keeping above mentioned points in mind , this is where fan in pattern excels.
	// * So we are no longer waiting for joe to speak.
	// * We are letting alice to communincate if he have something to say.
	// **/
	// c:=fanInUsingSelect(generatorPattern("Joe"),generatorPattern("Alice"))
	// for{
	// 	fmt.Println(<-c)
	// }

	// //--------SELECT PATTERN------------
	// c1:=generatorPattern("Joe")
	// c2:=generatorPattern("Alice")
	// for{
	// select {
	// case v1:=<-c1:
	// 	fmt.Printf("Received from channel 1 : %v",v1)
	// case v2:=<-c2:
	// 	fmt.Printf("Received from channel 2: %v",v2)
	// // default:
	// // 	fmt.Printf("No one was ready to communicate")
	// }

	// //-------------TIMEOUT SELECT CHANNEL---------------
	// /**
	// * So now we have this pattern , in which if channel doesn't return value for some time period.
	// * We will finish executing it.
	// * Here time.After returns current time after a certain period of time that is being passed as an param.
	// * So what happens is in each iteration we go inside switch statement.
	// * If we don't receive a value in first case we go to second case and start this timer.
	// * If we get the value from channel then we execute that case and we move to next iteration , time channel is null and void.
	// * But if we don't receive the value , we will receive value after one second in time channel.
	// * And we finish the loop.
	// **/
	// c:=generatorPattern("Ravi")
	// for{
	// 	select{
	// 	case v1:=<-c:fmt.Println(v1)
	// 	case <-time.After(time.Duration(1)*time.Second):return
	// 	}
	// }

	// //----------TIMEOUT FOR WHOLE CONVERSATION-----------
	// c:=generatorPattern("Ravi")
	// tc:=time.After(time.Duration(2)*time.Second)
	// for{
	// 	select{
	// 	case v1:=<-c:fmt.Println(v1)
	// 	case <-tc:fmt.Println("I am out")
	// 		return
	// 	}
	// }

	// //--------RECEIVE ON QUIT CHANNEL------------
	// /**
	// * In this we are telling go routine to end.
	// * And then before finishing that routine is communicating back.
	// * This is the pattern that is used alot in Redbelly.
	// **/
	// 	quit:=make(chan string)
	// 	communicateBack(generatorPattern("Ravi"),quit)
	// 	time.Sleep(time.Duration(1)*time.Second)
	// 	quit<-"Ohkay bie Ravi"
	// 	fmt.Println(<-quit)

	/**
	* --------DAISY CHAINING----------
	* It can be thought of as method of linking go routines that consumes and produce.
	**/
	c:=make(chan int)
	receiveChannel:=c
	var sendChannel chan int
	noOfRoutines:=10
	for i:=0;i<noOfRoutines;i++{
		sendChannel=make(chan int)
		go gopher(receiveChannel,sendChannel)
		receiveChannel=sendChannel
	}
	c<-1
	fmt.Println("Output from assembly line:",<-sendChannel)
}

func boring(msg string){
	for i:=0;;i++{
		fmt.Println(msg,i)
		time.Sleep(time.Duration(rand.Intn(1e3))*time.Millisecond)
	}
}

/**
* So in this generator pattern we are basically returning one side channel as return value.
* <-chan string means that channel returned from this function can only receive.
**/
func generatorPattern(msg string) <-chan string{
	c:=make(chan string)
	// The reason we have this function literal that is running as go routine.
	// Because we want to return channel from this function.
	// If there is no go routine then we will get stuck here only and won't return a value.
	go func(c chan string){
		for i:=0;;i++{
			c<-fmt.Sprintf(msg,i)
			// In this time.Duration expects time in nanoseconds , so time.Milliseconds basically means how many milliseconds in 1 nanosecond.
			// Also as per current understanding Sleep accepts time in nanoseconds.
			time.Sleep(time.Duration(rand.Intn(1e3))*time.Millisecond)
		}
	}(c)
	return c
}

func generatorPattern2(msg string) <-chan string{
	c:=make(chan string)
	// The reason we have this function literal that is running as go routine.
	// Because we want to return channel from this function.
	// If there is no go routine then we will get stuck here only and won't return a value.
	go func(c chan string){
		for i:=0;;i++{
			c<-fmt.Sprintf(msg,i)
			// In this time.Duration expects time in nanoseconds , so time.Milliseconds basically means how many milliseconds in 1 nanosecond.
			// Also as per current understanding Sleep accepts time in nanoseconds.
			// time.Sleep(time.Duration(rand.Intn(1e3))*time.Millisecond)
		}
	}(c)
	return c
}

func fanIn(input1,input2 <-chan string) <- chan string{
	c:=make(chan string)
	go func(){for{c <- <-input1}}()
	go func(){for{c <- <-input2}}()
	return c
}

// In this pattern in comparison to default fan-in pattern , we need only one go routine.
func fanInUsingSelect(input1,input2 <-chan string) <-chan string{
	/**
	* So for some reason var c chan string && c:=make(chan string) are behaving different.
	* So seems like , like slices , there is concept of nil and empty channels.
	**/
	c:=make(chan string)
	go func(){
		for{
			select{
			case v1:=<-input1:
				c<-v1
			case v2:=<-input2:
				c<-v2
			}
		}
	}()
	return c;
}

func communicateBack(talk <-chan string,quit chan string){
	go func(){
		for{
			select{ 
			case v1:=<-talk:fmt.Println(v1)
			case v2:=<-quit:fmt.Println(v2)
			quit<-"Ohkay bie!!"
				return
			}
		}
	}()
}



func gopher(receiveChannel chan int,sendChannel chan int){
	value:=<- receiveChannel
	sendChannel <- value+1
}


/**
* -------------SOME IMPORTANT POINTS REGARDING GO ROUTINES------------
* 1. That go routines and OS threads are not same thing. Go routine thread is handled by go scheduler and CPU might not even be aware of how many
* Go routines there are running in go lang env.
* 2. OS threads are more expensive in comparison to these go threads or go routines.
* 3. After main function execution finishes , go scheduler also goes down that's why even if there is any child go routine running , it will stop as it won't be
* scheduled by scheduler , that's why unless we use channels or sleep timer , main function will complete it's execution after invoking and we won't see any output
* from go routine.
* 4. We can see in generator pattern that even if parent function finishes , child routine will still persist given main function haven't finished.
**/
