# Library Management App with Go

## Project Overview

This Library Management App is a console-based application built with Go.  
It allows users to manage books and members through a terminal interface.

The application uses:

- Structs
- Interfaces
- Methods
- Slices
- Maps
- Error handling
- Package organization in Go

## Features

- Add a new book
- Remove an existing book
- Borrow a book
- Return a borrowed book
- List all available books
- List all borrowed books by a member

## Architecture

The application follows a simple layered architecture:

### Models

The `models` package contains the data structures used in the application.

- `Book` represents a library book with:
  - ID
  - Title
  - Author
  - Status

- `Member` represents a library member with:
  - ID
  - Name
  - List of borrowed books

### Services

The `services` package contains the business logic.

The `Library` struct manages:

- A map of books using book ID as the key
- A map of members using member ID as the key

The `LibraryManager` interface defines the operations supported by the library:

- Add books
- Remove books
- Borrow books
- Return books
- List available books
- List borrowed books

### Controllers

The `controllers` package handles communication with the user.

Responsibilities:

- Receive input from the terminal
- Validate user actions
- Call the appropriate service methods
- Display results to the user

### Main

The `main.go` file is the entry point of the application.

It:

- Creates the library service
- Creates the controller
- Runs the terminal menu
- Handles user choices

## How It Works

1. The user selects an operation from the terminal menu.
2. The controller receives the input.
3. The controller calls the corresponding library service method.
4. The service updates the library data.
5. The result is displayed back to the user.

## Error Handling

The application handles common errors such as:

- Trying to borrow a book that does not exist
- Trying to borrow an already borrowed book
- Trying to access a member that does not exist
- Returning a book that was not borrowed

## Data Storage

The application uses in-memory data storage:

- Books are stored in a map:

The data exists only while the program is running.

## Running the Application

From the project root:

```
go run main.go
```
