package main

import (
	"fmt"
)

func main(){
	fmt.Println("Slice in depth")
	/**
	* So in array , default value is 0 value of whatever data that array is off.
	* But incase of slice that 0 value can be thought of as nil slice.
	* This is because in array you specify number of elements it will store.
	* So if an array is of int type of size 5 then it will contain 5 zero's.
	* But incase of slice it is just a reference to an array. So if array doesn't exist then reference will be nil.
	**/
	var nilSlice []int
	fmt.Println("Value of uninitialized slice of type int:",nilSlice==nil," value of nilSlice:",nilSlice)

	/**
	* Slice is just an reference to array.
	* Then one may ask how is that possible when array is static and stores fix no. of elements.
	* Slice acts dynamic by creating another array of appropriate size when you want to add more elements.
	* And then referencing that array.
	* As shown below
	**/
	var fillArray =[5]int{1,2,3,4,5}
	nilSlice=fillArray[0:3]
	// Now you will see nilSlice won't be nil anymore
	fmt.Println("Value of uninitialized slice of type int:",nilSlice==nil," value of nilSlice:",nilSlice)
	//Now to show that nilSlice is referencing fillArray
	fillArray[0]=26
	fmt.Println("Value of uninitialized slice of type int:",nilSlice==nil," value of nilSlice:",nilSlice)

	/**
	* So in array we have len function to measure length of the array.
	* This len method works in the same in slices too and will return no. of elements slice is referencing to.
	* But we have another method by name cap.
	* This tells capacity of slice , which is number of elements from start index where it is referencing array.
	* Because this is the capacity that array can hold.
	**/
	fmt.Println("Length of nilSlice:",len(nilSlice)," capacity of slice:",cap(nilSlice))
	// 	You can't access 4th element just using index even though it is in underlying array.
	// For that you will create another slice and increase window using existing slice.
	fmt.Println("4th element of nilSlice :",nilSlice[0:5])

	/**
	* Slice is actually an struct , whose structure is as shown below.
	* type slice struct {
    	zerothElement *type
    	len int
    	cap int
	}
	* Here zerothElement is the pointer to the first element of the underlying array.
	* So if we are initializing it by a[2:6] then 2nd element address is zerothElement value.
	* len is the length of the slice.
	* cap is the capacity.
	* That's why when slice is uninitialized then zerothElement is set to nil.
	**/
	fmt.Println("Address in actual array:",&fillArray[0])
	fmt.Println("Address in slice:",&nilSlice[0])
	/**
	* append function is used append new elements in slice.
	* It is variadic function.
	* It creates a new array to store more elements in slice.
	**/
	// So here since we are adding 1 element which is still under the capacity of slice.
	// So underlying array wont change and addresses will be same.
	// This is quite confusing here it won't really append.
	// But create new slice with nilSlice[0:4] and set nilSlice[3]=10.
	// As a result value of underlying array will also change.
	// So this tells when you initialize slice like this , slice acts more like a window to underlying array.
	newSlice:=append(nilSlice,10)
	fmt.Println("This is the new slice:",newSlice)
	fmt.Println("This is the old array:",fillArray)
	fmt.Println("Address in actual array:",&fillArray[0])
	fmt.Println("Address in new slice:",&newSlice[0])
	// Here we are adding more elements then capacity so it will create new array and addresses will be different.
	newSlice2:=append(nilSlice,10,11,12,13,14,15)
	fmt.Println("This is the new slice:",newSlice2)
	fmt.Println("Address in actual array:",&fillArray[0])
	fmt.Println("Address in new slice2:",&newSlice2[0])
	/**
	* So another way to create an slice is to let go create underlying anonymous array.
	**/
	newSlice3:=[]int{1,2,3}
	fmt.Println("So this is new slice created using anonymous function:",newSlice3," length:",len(newSlice3)," capacity:",cap(newSlice3))
	newSlice4:=append(newSlice3,56)
	// So it seems like, while creating a new array , it just doubles the length of initial array.
	fmt.Println("So this is new slice created using anonymous function and appending to it:",newSlice4," length:",len(newSlice4)," capacity:",cap(newSlice4))
	/**
	* So in above examples , we learned about nil slice.
	* We know that nil slice is basically an slice which doesn't have any underlying array and zerothElement is set to nil.
	* Empty slice is slice which has an array which is empty.
	* In go make function allows you to create this empty slice.
	**/
	sliceUsingMake:=make([]int,2,4)
	fmt.Println("This is slice we made using make:",sliceUsingMake," This is the length:",len(sliceUsingMake)," This is the capacity:",cap(sliceUsingMake))

	/**
	* One thing to keep in mind is that in comparison to common assumption.
	* Slices are actually passed by value in go lang.
	* But since they refer to same underlying array , it seems like they are being passed by reference.
	**/

	/**
	* -----Slice comparison------
	* So go lang says that you can only compare or check slice for nil condition.
	* To check a non-nil slice you will have to use for loop.
	**/
}