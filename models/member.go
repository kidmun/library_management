package models

type Member struct {
	ID            int `json:"id"`
	Name          string `json:"name"`
	BorrowedBooks []Book `json:"borrowed_books"`
}