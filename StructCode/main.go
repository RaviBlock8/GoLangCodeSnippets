package main

import "fmt"

// So here we are creating a struct

type person struct{
	firstName string
	lastName string
}
type business struct{
	companyName string
}

type buisnessMan struct{
	personDetails person
	businessDetails business
}
/**
* There are multiple ways to initialize struct as shown below
* ravi:=person{"Ravi","Verma"}
* nitu:=person{firstName:"Nitu",lastName:"thakur"}
**/
 
func main(){
	ravi:=person{firstName: "Ravi",lastName: "Verma"}
	raviVerma:=buisnessMan{personDetails: person{firstName: "Ravi",lastName: "Verma"},businessDetails: business{companyName: "Block8"}}
	fmt.Println(ravi)
	fmt.Println(raviVerma)
	/**
	* So we have this concept of pointers in go lang.
	* Essentially they are used because when we call receiver function.
	* We want to apply changes directly to data or variable in reference to which we are calling these functions.
	* Without using pointer , it will pass a copy of ravi to receiver method.
	* One way to call these functions in reference to pointer is to do it in traditional way as shown below.
	**/
	// pointerToRavi:=&ravi
	// pointerToRavi.updateName("Nitu")
	// fmt.Println(ravi)
	// Another short way to do soo , is to not to use & , and go will convert it directly to an pointer by looking at argument type in receiver function.
	ravi.updateName("Ramesh")
	fmt.Println(ravi)
	testSlice:=[]string{"Ravi","verma"}
	updateSlice(testSlice)
	fmt.Println(testSlice)
}

func (pointerToPerson *person) updateName(newName string){
	(pointerToPerson).firstName=newName
}

func updateSlice(arg []string){
	arg[0]="changed!!"
}

/**
*---------------SLICE-----------------------------
* So important thing to note about slice is how this slice is created in memory.
* So when you create a slice , it has two parts.
* One is the underlying array that stores all the elements.
* Then there is actual slice structure , which contains 3 elements.
* 1. length of slice 2. Capacity of slice 3. Pointer to first element
* So when you pass this slice as an argument , go creates an copy of slice data structure.
* But underlying array still stays the same
**/