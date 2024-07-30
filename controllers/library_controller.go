package controllers

import (
	"encoding/json"
	"fmt"
	"library_management/models"
	"library_management/services"
	"net/http"
	"strconv"
)

var library = services.NewLibrary()
func AddBookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }
	var book models.Book
	if err:= json.NewDecoder(r.Body).Decode(&book); err != nil{

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	library.AddBook(book)
	fmt.Fprintf(w, "Book added: %v", book)
}
func ListAvailableBooksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }
	availableBooks := library.ListAvailableBooks()
	json.NewEncoder(w).Encode(availableBooks)

}
func RemoveBookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }
	bookID, err := strconv.Atoi(r.URL.Query().Get("id"))
    if err != nil {
        http.Error(w, "Invalid book ID", http.StatusBadRequest)
        return
    }
	err = library.RemoveBook(bookID)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Book Deleted")
}

func BorrowBookHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPut {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }
	bookID, err := strconv.Atoi(r.URL.Query().Get("bookId"))
    if err != nil {
        http.Error(w, "Invalid book ID", http.StatusBadRequest)
        return
    }
	memberId, err := strconv.Atoi(r.URL.Query().Get("memberId"))
    if err != nil {
        http.Error(w, "Invalid member ID", http.StatusBadRequest)
        return
    }

	err = library.BorrowBook(bookID, memberId)
	if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	fmt.Fprintf(w, "Book with ID %d borrowed by member with ID %d", bookID, memberId)
}

func ReturnBookHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPut {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }
	bookID, err := strconv.Atoi(r.URL.Query().Get("bookId"))
    if err != nil {
        http.Error(w, "Invalid book ID", http.StatusBadRequest)
        return
    }
	memberId, err := strconv.Atoi(r.URL.Query().Get("memberId"))
    if err != nil {
        http.Error(w, "Invalid member ID", http.StatusBadRequest)
        return
    }
	err = library.ReturnBook(bookID, memberId)
	if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	fmt.Fprintf(w, "Book with ID %d returned by member with ID %d", bookID, memberId)
}

func ListBorrowedBooksHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }
	memberId, err := strconv.Atoi(r.URL.Query().Get("id"))
    if err != nil {
        http.Error(w, "Invalid member ID", http.StatusBadRequest)
        return
    }

	books, err := library.ListBorrowedBooks(memberId)

	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(books)
}

func AddMemberHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }
	var member models.Member
	if err:= json.NewDecoder(r.Body).Decode(&member); err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	member.BorrowedBooks = []models.Book{}
	library.AddMember(member)
	fmt.Fprintf(w, "Member added: %v", member)
}