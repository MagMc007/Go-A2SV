package main

import (
	"fmt"

	"library_management/controllers"
	"library_management/services"
)


func main() {
	library := services.NewLibrary()
	controller := controllers.NewLibraryController(library)

	var choice int

	for {
		fmt.Println(`Hello there what would You like(use the number)  
			1. Add a new book        
			2. Remove an existing        
			3. Borrow a book         
			4. Return a book         
			5. List all available books         
			6. List all borrowed books     
			7. Terminate
		`)

		fmt.Scan(&choice)

		switch choice {
			case 1:
				controller.AddBook()
			case 2:
				controller.RemoveBook()
			case 3:
				controller.BorrowBook()
			case 4:
				controller.ReturnBook()
			case 5:
				controller.ListAvailableBooks()
			case 6:
				controller.ListBorrowedBooks()
			case 7:
				fmt.Println("Good Bye")
				return
			}
	}

}