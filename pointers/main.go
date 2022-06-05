package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println("Let's learn about pointers")
	var num int = 26
	list := []int{1, 2, 23}
	ravi := person{name: "ravi", age: 25}
	fmt.Printf("Integer value:%v\n", num)
	fmt.Printf("struct value:%v\n", ravi)
	pnum := &num
	pravi := &ravi
	fmt.Printf("Integer pointer value:%v\n", *pnum)
	fmt.Printf("struct pointer value:%v %v\n", pravi.name, pravi.age)
	fmt.Printf("Value of list:%v\n", list)
	fmt.Printf("Address of list:%v\n", &list)
}
