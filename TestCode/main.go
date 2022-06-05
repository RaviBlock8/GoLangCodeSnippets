package main

import "fmt"

func main() {
	list := []int{1, 2, 3, 4, 5, 6}
	var n uint32
	sendChan := make(chan uint32)
	receiveChan := make(chan []int)

	go func(l []int) {
		n = <-sendChan
		receiveChan <- l[0:n]
	}(list)

	sendChan <- 2
	fmt.Println(<-receiveChan)
}
