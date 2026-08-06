package main

type LibraryManager interface {
	AddBook(book Book)
	RemoveBook(bookID int)

	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error

	ListAvailableBooks() []Book
	ListBorrowedBooks(memeberID int) []Book
}

