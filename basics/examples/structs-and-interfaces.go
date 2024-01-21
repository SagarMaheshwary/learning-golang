package examples

import (
	"fmt"
)

func StructsAndInterfaces() {
	//STRUCTS

	//Defining a struct
	type User struct {
		ID    int
		Name  string
		Email string
	}

	//Instantiating a struct
	user := User{
		ID:    1,
		Name:  "John",
		Email: "john@gmail.com",
	}

	//Reading a field
	fmt.Println("User Name", user.Name)

	//NESTED STRUCTS

	type Post struct {
		ID     int
		Title  string
		Author User
	}

	post := Post{
		ID:     1,
		Title:  "Post One",
		Author: user,
	}

	fmt.Println("Post Author", post.Author.Name)

	//EMBEDDED STRUCTS

	type CustomerDetail struct {
		Address string
		Phone   string
	}

	type Customer struct {
		ID   int
		Name string
		CustomerDetail
	}

	customer := Customer{
		ID:   1,
		Name: "John",
		CustomerDetail: CustomerDetail{
			Address: "Test",
			Phone:   "12345",
		},
	}

	//Embedded structs can be accessed as if they are defined on the main struct.

	fmt.Println("Customer Address", customer.Address)
	fmt.Println("Customer Phone", customer.Phone)

	fmt.Println(user, post)
}
