package services
import (
	"errors"

	"library_management/models"
)


type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int)

	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error

	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
}


// declare the struct
type Library struct {
	bookMap map[int]models.Book
	memberMap map[int]models.Member
}

func NewLibrary() *Library {
	return &Library{
		bookMap: make(map[int]models.Book),
		memberMap: make(map[int]models.Member),
	}
}

// method implementations
func (l *Library) AddBook(book models.Book) {
	l.bookMap[book.ID] = book
}

func (l *Library) RemoveBook(bookID int) {
	delete(l.bookMap,bookID)
}

func (l *Library) BorrowBook(bookID int, memberID int) error {
	// is book avalable 
	book, exists := l.bookMap[bookID]

	if !exists {
		return errors.New("Book does not exist")
	}

	// check the book status
	if book.Status != "Available" {
		return errors.New("Book has been borrowed")
	}
	// does member exist
	member, exists :=  l.memberMap[memberID]

	if !exists {
		return errors.New("Member does not exist")
	}


	member.BorrowedBooks = append(member.BorrowedBooks, book)

	book.Status = "Borrowed"

	// update the map with the newly updated memeber and book
	l.bookMap[bookID] = book
	l.memberMap[memberID] = member

	return nil
}


func (l *Library) ReturnBook(bookID int, memberID int) error{
	// is book avalable 
	book, exists := l.bookMap[bookID]

	if !exists {
		return errors.New("Book does not exist")
	}

	// check the book status
	if book.Status != "Borrowed" {
		return errors.New("Book has was not Borrowed")
	}

	// does member exist
	member, exists :=  l.memberMap[memberID]

	if !exists {
		return errors.New("Member does not exist")
	}

	// make the book availbale
	book.Status = "Available"

	// remove it from the memebers borrowed slice
	for i, borrowedBook := range member.BorrowedBooks {
		if borrowedBook.ID == bookID {
			member.BorrowedBooks = append(
				member.BorrowedBooks[:i],
				member.BorrowedBooks[i+1:]...,
			)
			break
		}
	}


	// update the map
	l.bookMap[bookID] = book
	l.memberMap[memberID] = member

	return nil
} 

func (l *Library) ListAvailableBooks() []models.Book {
	// prepare an array append each to it from the map
	books := make([]models.Book, 0)

	for _, book := range l.bookMap {
		if book.Status == "Available" {
			books = append(books, book)
		}
	}

	return books
}

func (l *Library) ListBorrowedBooks(memberID int)  []models.Book {
	books := make([]models.Book, 0)
	// does the member exist
	member, exists := l.memberMap[memberID]

	if !exists {
		return books
	}

	return member.BorrowedBooks
}