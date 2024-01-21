package examples

import (
	"fmt"
	"slices"
)

func ArraysAndSlices() {
	//ARRAYS arrays in golang are fixed sized.
	//Since they are fixed size so all the array values are stored sequentially in the memory.

	var names [3]string //declaring an array of string values.

	fmt.Println("Names at 0 index", names[0]) //reading a value by index.

	names[0] = "Daniel" //assigning a value
	names[1] = "John"
	names[2] = "Jane"

	ages := [3]int{1, 2, 3}       //Initializing values
	moreAges := [...]int{1, 2, 3} //we can also let the compiler infer the array length with "..."

	//SLICES, An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible
	//view into the elements of an array. In practice, slices are much more common than arrays.

	//Declaring slice is similar to an array expect for defining a size.
	var stocks []int

	prices := []int{1, 2, 3} //Initializing a slice
	//NOTE: When adding an element to the slice, the underlying array will be assigned a new memory location.
	prices = append(prices, 4)           //Adding an element to a slice
	prices = slices.Delete(prices, 0, 1) //Deletes "1" from the slice
	prices = slices.Delete(prices, 0, 2) //Deletes "2","3" from the slice

	//We can also create slices with make() builtin function
	//make(type, size, capacity), if no capacity is passed then size will be used as capacity.
	websites := make([]string, 3, 5)

	//To find the size of a slice, we can use len() builtin function.
	totalWebsites := len(websites)

	//And to find the capacity of a slice, we can use cap() builtin function.
	websitesToStore := cap(websites)

	//Creating a slice from an existing array
	males := names[0:2]   //From start to second element, prints: ["Daniel", "John"]
	myNames := names[:]   //All elemnts, prints: ["Daniel", "John", "Jane"]
	myNames1 := names[1:] //All elements after 1 index, prints: ["John", "Jane"]
	myNames2 := names[:1] //All elements till 1 index, prints: ["Daniel"]

	//-------------------------------------------------------------------------------------------------------------------------------------------
	//Ignore this line as it's only written for compiler to silent it's "unused vars" error
	fmt.Println(ages, moreAges, stocks, prices, websites, totalWebsites, websitesToStore, males, myNames, myNames1, myNames2)
}
