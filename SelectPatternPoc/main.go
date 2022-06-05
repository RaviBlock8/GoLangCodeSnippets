package main

import (
	"fmt"
	"time"
)

func handleProposal(proposalChan chan int) {
	for {
		select {
		case proposal := <-proposalChan:
			fmt.Printf("Got proposal no: %v and now we will start doing some sync work\n", proposal)
			time.Sleep(time.Duration(5) * time.Second)
			fmt.Printf("Done doing sync work for proposal no : %v\n", proposal)
		}
	}
}
func main() {
	proposalChan := make(chan int)
	go handleProposal(proposalChan)
	for i := 0; i <= 5; i++ {
		proposalChan <- i
	}
}
