package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Lets go......")
	bn := uint64(12344)
	sbn := strconv.FormatUint(bn, 10)
	fmt.Printf("String conversion : %v", sbn)
	val, err := strconv.ParseUint(sbn, 10, 64)
	if err != nil {
		fmt.Println("cant parse")
		return
	}
	fmt.Printf("String conversion to uint : %v", val)
}
