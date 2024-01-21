package main

import (
	"fmt"
	"strconv"
)

func main() {
	//STRING

	//Declaring a variable
	var someData string

	//Assigning values
	someData = "whatever"

	//Initializing a variable
	var name string = "John Doe"

	//Shorthand for initializating a variable where type will be inferred
	anotherName := "Jane Doe"

	//Constants (cannot reassign value afterwards, can only be declared with const keyword)
	const Os = "Linux"

	//Declaring multiple variables (NOTE: with var you can only declare/initialize same types)
	var var1, var2 string

	//Initializing multiple variables
	var var3, var4 string = "var1", "var2"

	//Or using shorthand syntax
	var5, var6 := 8, true // var5= integer, var6=boolean

	//BOOLEAN

	var isGoGood = true

	//INTEGERS

	//int is an alias to int32/int64 depending on your CPU architecture
	var age int8 = 25             //8 bit/1 byte integer, min=-128 max=127
	var quantity int16 = 32767    //16 bit/2 byte integer, min=-32768 max=32767
	var amount int32 = 1000       //32 bit/4 byte integer, min=-2147483648 max=2147483647
	var revenue int64 = 100000000 //64 bit/8 byte integer, min=-9223372036854775808 max=9223372036854775807

	//We also have unsigned integers that can only store positive values
	//Types: uint8 uint16 uint32 uint64 uint
	//uint8 max=255
	//uint16 max=65535
	//uint32 max=4294967295
	//uint64 max=18446744073709551615

	//FLOATS

	var price float32 = 999.99    //min=-3.4e+38 max= 3.4e+38
	var total float64 = 999999.99 //min=-1.7e+308 max=+1.7e+308

	//COMPLEX, takes 2 floating point literals. one for the real part and one for the imaginary part

	var complexNumber complex64 = complex(10, 11)
	var anotherComplexNumber complex128 = complex(128, 129)

	//BYTE alias for int16, used for representing ASCII characters

	var byteCode byte = 65
	var byteChar byte = 'A' // prints: 65, as it's the character code for "A"

	//RUNE alias for int32, used for representing unicode code point

	var runeCode rune = 97
	var runeChar rune = 'a' // prints: 97, as it's the character code for "a"

	//One thing to note in Golang is that if we're initializing a value that's above the range of a type then
	//you will get error from compiler but if we're changing it afterwards then it won't, rather it will create
	//unexpected results. e.g quantity variable has been assigned the maximum value an int16 can hold so if we do:
	// quantity += 1 // we'll get: -32767

	//Default VALUES WHEN DECLARING A VARIABLE
	//String: ""
	//Boolean: false
	//All Integer types: 0

	//TYPECASTING
	//Since Golang is strictly typed language so we can't do operations on different types of variables

	sum := price + float32(quantity) //Here we need to convert integer (int16) to float before addition operation
	intPrice := int(price)           //float to int
	sum = float32(10 + uint(price))  //we also can't do int + uint so we first need to convert the variable to unit

	//We can convert float to string using the fmt package
	floatToString := fmt.Sprintf("%f", price)

	//Or using strconv package, also note that the first argument takes float64 (faster than fmt.Sprintf)
	floatToString2 := strconv.FormatFloat(float64(price), 'E', -1, 32)

	//Converting an int to a stirng like this will result in a Rune value
	str := string(int(sum)) // prints "ϱ"

	//Here we can use the strconv package to property convert it to string
	str = strconv.FormatInt(int64(sum), 2)

	//Or with fmt.Sprintf
	str = fmt.Sprintf("%d", int(sum))

	//-------------------------------------------------------------------------------------------------------------------------------------------
	//Ignore this line as it's only written for compiler to silent it's "unused vars" error
	println(Os, someData, var1, var2, var3, var4, var5, var6, isGoGood, name, anotherName, age, quantity, amount, revenue, price, total, complexNumber, anotherComplexNumber, runeCode, runeChar, byteCode, byteChar, intPrice, str, floatToString, floatToString2)
}
