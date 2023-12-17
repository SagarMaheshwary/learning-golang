package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Book struct {
	Id     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var books []Book

func main() {
	fmt.Println("Starting Application")

	books = append(books, Book{Id: uuid.NewString(), Isbn: "1234", Title: "Book one", Author: &Author{FirstName: "Daniel", LastName: "John"}})
	books = append(books, Book{Id: uuid.NewString(), Isbn: "1234", Title: "Book Two", Author: &Author{FirstName: "Daniel", LastName: "John"}})

	router := mux.NewRouter()

	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func getBooks(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	json.NewEncoder(writer).Encode(books)
}

func getBook(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)

	for _, book := range books {
		if book.Id == params["id"] {
			json.NewEncoder(writer).Encode(book)
			return
		}
	}

	json.NewEncoder(writer).Encode("Not Found")
}

func createBook(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var book Book

	_ = json.NewDecoder(req.Body).Decode(&book)
	book.Id = uuid.NewString()
	books = append(books, book)

	json.NewEncoder(writer).Encode(book)
}

func updateBook(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)

	for index, book := range books {
		if book.Id == params["id"] {
			var newBook Book
			json.NewDecoder(req.Body).Decode(&newBook)

			(&books[index]).Isbn = newBook.Isbn
			(&books[index]).Title = newBook.Title

			break
		}
	}

	json.NewEncoder(writer).Encode(books)
}

func deleteBook(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)

	for index, book := range books {
		if book.Id == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}

	json.NewEncoder(writer).Encode(books)
}
