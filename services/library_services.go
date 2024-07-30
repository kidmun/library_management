package services

import (
	"errors"
	"library_management/models"
)

type LibraryManager interface {
	AddBook(book models.Book) 
	ListAvailableBooks() []models.Book
	RemoveBook(bookId int)
	BorrowBook(bookId, memberId int) error
	ReturnBook(bookId, memberId int) error
	ListBorrowedBooks(memberId int) ([]models.Book, error)
	AddMember(member models.Member)
}

type Library struct{

	Books map[int]models.Book
	Members map[int]models.Member
}

func NewLibrary()*Library{
	return &Library{
		make(map[int]models.Book),
		make(map[int]models.Member),
	}
}
func (lib *Library) AddBook(book models.Book){
	lib.Books[book.ID] = book
}

func (lib *Library) ListAvailableBooks() []models.Book{

	books := []models.Book{}
	
	for _, val:= range lib.Books{
		books = append(books, val)
	}
	return books
}

func (lib *Library) RemoveBook(bookId int) error{
	for id, _:= range lib.Books{
		if id == bookId{
			delete(lib.Books, bookId)

			return nil
		}
	}
	return errors.New("book not found")
	
}

func (lib *Library) BorrowBook(bookId, memberId int) error{


	book, bookExists := lib.Books[bookId]
	if !bookExists{
		return errors.New("book not found")
	}
	member, memberExists := lib.Members[memberId]
	if !memberExists{
		return errors.New("member not found")
	}

	if book.Status == "Borrowed"{
		return errors.New("book already borrowed")
	}
	book.Status = "Borrowed"
	lib.Books[bookId] = book
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	lib.Members[memberId] = member
	return nil
}

func (lib *Library) ReturnBook(bookId, memberId int) error{

	book, bookExists := lib.Books[bookId]
	if !bookExists{
		return errors.New("book not found")
	}
	member, memberExists := lib.Members[memberId]
	if !memberExists{
		return errors.New("member not found")
	}
	book.Status = "Available"
	lib.Books[bookId] = book
	for i, cur := range member.BorrowedBooks{
		if cur.ID == bookId{
			member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i + 1:]...)
			break
		}
	}
	return nil

}

func (lib * Library)ListBorrowedBooks(memberId int)([]models.Book, error){
	member, memberExists := lib.Members[memberId]
	if !memberExists{
		return nil, errors.New("member not found")
	}
	return member.BorrowedBooks, nil
}
func (lib *Library) AddMember(member models.Member){
	lib.Members[member.ID] = member
}