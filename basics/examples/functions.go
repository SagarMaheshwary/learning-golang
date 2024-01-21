package examples

import (
	"errors"
	"fmt"
)

//Note: functions that exported (used in other file/package) as defined using pascal-case
//and internal functions with camel-case

// Basic function definition
func Functions() {
	exampleFunc() //Calling a function

	sum := sum(1, 2)
	substractedVal := substract(1, 2)

	result, err := divide(1, 1)

	if err != nil {
		fmt.Println("Error", err)
	}

	//We can use "_" to ignore a variable that we don't need to use
	result1, _ := divide(1, 1)

	num := 1

	//Variables in Golang are passed by values so we can use pointers
	//to pass and modify the original value inside the function.
	increment(&num)

	firstName, lastName := getNames()
	fullName := getFullName()

	//-------------------------------------------------------------------------------------------------------------------------------------------
	//Ignore this line as it's only written for compiler to silent it's "unused vars" error
	fmt.Println(sum, substractedVal, result, result1, firstName, lastName, fullName)
}

func exampleFunc() {
	fmt.Println("Example Function.")
}

// With Params and returning value
func sum(a int, b int) int {
	return a + b
}

// If all the params have same type then we only need to defined it at the end
func substract(a, b int) int {
	return a - b
}

// Returning multiple values
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}

	result := a / b

	return result, nil
}

// Using pointer as the param
func increment(num *int) {
	*num++
}

// We also have the ability to create named returns but they are generally not preferred due to less readability:
func getNames() (firstName string, lastName string) {
	firstName = "John"
	lastName = "Doe"

	return //we just need to call the return statement without needing to specify the return params
}

func getFullName() (name string) {
	name = "John Doe"

	return "Jane Doe" //explicit return overrides the named returns
}
