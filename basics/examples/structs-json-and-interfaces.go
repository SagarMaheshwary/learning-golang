package examples

import (
	"encoding/json"
	"fmt"
)

func StructsJsonAndInterfaces() {
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
		Address, Phone string //Multiple fields with same type can also be defined like this
	}

	type Customer struct {
		ID   int
		Name string
		CustomerDetail
	}

	//Instantiating an embedded struct
	customer := Customer{
		ID:   1,
		Name: "John",
		CustomerDetail: CustomerDetail{
			Address: "Test",
			Phone:   "12345",
		},
	}

	//Although embedded structs are instantiated like nested structs but they can be accessed
	//as if they are defined on the main struct.
	fmt.Println("Customer Address", customer.Address)
	fmt.Println("Customer Phone", customer.Phone)

	//CONVERTING STRUCTS TO JSON OR VICE VERSA
	type Article struct {
		ID          int    `json:"id"` //these are the keys used when converting to json
		Title       string `json:"title"`
		PublishedAt string `json:"published_at"`
	}

	article := Article{
		ID:          1,
		Title:       "Article One",
		PublishedAt: "Today",
	}

	//json.Marshal returns two values bytes array (we can convert to string) and error value in case it failed
	jsonUser, err := json.Marshal(&article)

	if err != nil {
		fmt.Println("Unable to create json.")
	}

	fmt.Println(string(jsonUser)) // Prints: {"id":1,"title":"Article One","published_at":"Today"}

	//Same as json.Marshal but the output is prettified
	//json.MarshalIndent(&article, "", " ")

	articleAsJson := `
		{
			"id": 2,
			"title": "Article Two",
			"published_at": "Today"
		}`

	article2 := new(Article)

	//We can use json.Unmarshal to parse json into a struct
	if err := json.Unmarshal([]byte(articleAsJson), article2); err != nil {
		fmt.Println("Unable to parse json.")
	}

	fmt.Println(article2)

	//STRUCT METHODS

	//In Go, you can associate methods with a struct by declaring the method with a receiver. A receiver is a
	//special parameter that allows a function to be associated with a specific type, including struct types.
	//NOTE: you can't define struct methods inside of a function so we have defined "Image" struct and "imageUrl"
	//method outside the function.

	image := Image{
		ID:   1,
		Path: "test-path",
	}

	fmt.Println(image.imageUrl())
}

// STRUCT METHODS

type Image struct {
	ID   int
	Path string
}

func (i Image) imageUrl() string {
	return fmt.Sprintf("http://localhost/images/%v", i.Path)
}
