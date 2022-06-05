package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	list := make([]string, 4, 4)
	list = append(list, "localhost:8545")
	list = append(list, "localhost:6443")
	fmt.Println(list)
	marshalledList, err := json.Marshal(list)
	if err != nil {
		log.Fatalf("Can't be marshalled")
	}
	fmt.Printf("\nMarshalled list: %v\n", marshalledList)
	encodedList := base64.StdEncoding.EncodeToString(marshalledList)
	fmt.Printf("Encoded list: %v", encodedList)
	fmt.Println("Let's get the list back")
	decodedList, err1 := base64.StdEncoding.DecodeString(encodedList)
	if err1 != nil {
		log.Fatalf("List can't be decoded")
	}
	var unmarshalledList []string
	err2 := json.Unmarshal(decodedList, &unmarshalledList)
	if err2 != nil {
		log.Fatalf("List can't be unmarshalled")
	}
	fmt.Printf("Here's the unmarshalled list:%v\n", unmarshalledList)
}
