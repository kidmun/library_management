package main

import (

	"library_management/controllers"
	"net/http"
)

func main() {
	
	http.HandleFunc("/books", controllers.ListAvailableBooksHandler)
	http.HandleFunc("/addBook", controllers.AddBookHandler)
	http.HandleFunc("/borrowBook", controllers.BorrowBookHandler)
	http.HandleFunc("/removeBook", controllers.RemoveBookHandler)
	http.HandleFunc("/returnBook", controllers.ReturnBookHandler)
	http.HandleFunc("/listBorrowedBooks", controllers.ListBorrowedBooksHandler)
	http.HandleFunc("/addMember", controllers.AddMemberHandler)
	http.ListenAndServe(":8080", nil)
}