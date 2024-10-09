package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"tserv/model"

	"github.com/gorilla/mux"
)

func main() {
	mux := mux.NewRouter()

	model.Books = append(model.Books, model.Book{ID: "1", Title: "Go developer", Autor: &model.Autor{FirstName: "John", LastName: "Gorilla"}})

	mux.HandleFunc("/books", getBooks).Methods("GET")
	mux.HandleFunc("/books/{id}", getBook).Methods("GET")
	mux.HandleFunc("/books", createBook).Methods("POST")
	mux.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	mux.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	fmt.Println("Starting...")

	log.Fatal(http.ListenAndServe(":8181", mux))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(model.Books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p := mux.Vars(r)
	for _, item := range model.Books {
		if item.ID == p["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(model.Books)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book model.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000))
	model.Books = append(model.Books, book)
	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range model.Books {
		if item.ID == params["id"] {
			model.Books = append(model.Books[:index], model.Books[index+1:]...)
			var book model.Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			model.Books = append(model.Books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(model.Books)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range model.Books {
		if item.ID == params["id"] {
			model.Books = append(model.Books[:index], model.Books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(model.Books)
}
