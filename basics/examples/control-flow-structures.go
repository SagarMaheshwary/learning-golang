package examples

import (
	"fmt"
	"strconv"
)

func ControlFlowStructures() {
	//Basic IF STATEMENT

	isProductAvailable := true

	if isProductAvailable { // Same as isProductAvailable == true
		fmt.Println("Yes Product is available.")
	}

	//IF-ELSE STATEMENT

	if !isProductAvailable { // Same as isProductAvailable != true
		fmt.Println("No Product is not available.")
	} else {
		fmt.Println("Yes Product is available.")
	}

	//IF-ELSE-IF STATEMENT

	quantity := 99

	if quantity >= 100 {
		fmt.Println("Quantity is either equals to or above 100.")
	} else if quantity >= 50 {
		fmt.Println("Quantity is either equals to or above 50.")
	} else {
		fmt.Println("Qantity is less than 50.")
	}

	//We can also initialize a variable and check through if statement
	//and that variable will only be available till the if block

	if val, err := strconv.ParseBool("true"); err != nil {
		fmt.Println("Coundn't parse the string.")
	} else {
		fmt.Println("String parsed succesfully", val)
	}

	//Here both val and err won't be available because their scope was
	//limited to if statement.

	//SWITCH-CASE

	//Golang only executes the selected case so there's no need for break statement
	switch quantity {
	case 10:
		fmt.Println("Quantity is 10.")
	case 20, 30: //You can pass multiple conditions to a single case.
		fmt.Println("Quantity is 20.")
	default:
		fmt.Println("Quantity is above 20.")
	}

	//We can also use switch-case like an if-else statement with different conditions

	age := 25

	switch {
	case age >= 12 && age < 18:
		fmt.Println("You're still a kid.")
	case age >= 18 && age < 25, age >= 25 && age < 40: //Here we have defined multiple conditions in one case statement
		fmt.Println("You're an adult now.")
	default:
		fmt.Println("You're getting old.")
	}

	//CONDITIONAL OPERATORS
	//==, equals
	//!=, not equal
	//>=, greater than or equal
	//<=, less than or equal
	//>, greater than
	//<, less than

	//LOGICAL OPERATORS
	//&&, and
	//||, or
	//!, not
}
