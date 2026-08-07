package controllers

import (
	"fmt"

	"library_management/services"
	"library_management/models"
)

type LibraryController struct {
	library *services.Library
}

func NewLibraryController(library *services.Library) *LibraryController {
	return &LibraryController{
		library: library,
	}
}

// add book
func (c *LibraryController) AddBook() {
	var id int
	var title string
	var author string

	fmt.Print("Enter book ID: ")
	fmt.Scan(&id)

	fmt.Print("Enter book title: ")
	fmt.Scan(&title)

	fmt.Print("Enter author: ")
	fmt.Scan(&author)

	book := models.Book{
		ID:     id,
		Title:  title,
		Author: author,
		Status: "Available",
	}

	c.library.AddBook(book)

	fmt.Println("Book added successfully")
}

func (c *LibraryController) RemoveBook() {
	var id int

	fmt.Print("Enter book ID: ")
	fmt.Scan(&id)


	c.library.RemoveBook(id)

	fmt.Println("Book removed successfully")
}


func (c *LibraryController) BorrowBook() {
	var bookID int
	var memberID int

	fmt.Print("Enter book ID: ")
	fmt.Scan(&bookID)

	fmt.Print("Enter memberID ")
	fmt.Scan(&memberID)

	err := c.library.BorrowBook(bookID, memberID)

	if err != nil {
		fmt.Println(err)
		return  
	}
	fmt.Println("Book borrowed successfully")
}


func (c *LibraryController) ReturnBook() {
	var bookID int
	var memberID int

	fmt.Print("Enter book ID: ")
	fmt.Scan(&bookID)

	fmt.Print("Enter memberID ")
	fmt.Scan(&memberID)

	err := c.library.ReturnBook(bookID, memberID)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Book returned successfully")
}


func (c *LibraryController) ListAvailableBooks() {

	books := c.library.ListAvailableBooks()

	fmt.Println(books)
}


func (c *LibraryController) ListBorrowedBooks() {
	var memberID int

	fmt.Println("Enter memberID:")
	fmt.Scan(&memberID)

	books := c.library.ListBorrowedBooks(memberID)

	fmt.Println(books)
}