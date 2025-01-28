package service

import (
	"fmt"
	"log"

	pb "github.com/BlazeCode1/book-grpc/app/book/controller/grpc"
	b "github.com/BlazeCode1/book-grpc/app/book/model/Book"
	operation "github.com/BlazeCode1/book-grpc/app/book/repository"
	"github.com/google/uuid"
)

func HandleGetBooks() (*pb.BookListResponse, error) {
	// remove this query and add it to repository/operations.go # DONE
	books, err := operation.GetBooks()
	if err != nil {
		return nil, fmt.Errorf("failed to get books: %v", err)
	}

	var pbBooks []*pb.Book
	for _, book := range books {
		pbBooks = append(pbBooks, &pb.Book{
			Id:       book.ID,
			BookName: book.BookName,
		})
	}
	return &pb.BookListResponse{Books: pbBooks}, nil
}

// here we require packacge model book # DONE
func HandleAddBook(book b.Book) (*pb.BookResponse, error) {
	log.Printf("Adding book: %s", book.BookName)

	// Create a new Book instance
	bookInstance := b.Book{
		ID:       uuid.New().String(), // Generate a new UUID for each book
		BookName: book.BookName,
	}

	err := operation.InsertBook(bookInstance)
	if err != nil {
		return nil, fmt.Errorf("failed to add book: %v", err)
	}

	return &pb.BookResponse{
		Message: fmt.Sprintf("Book '%s' added successfully", bookInstance.BookName),
	}, nil
}

func HandleDeleteBook(id string) (*pb.BookResponse, error) {
	log.Printf("Deleting book with ID: %s", id)
	// Delete the book from your storage
	err := operation.DeleteBook(id)
	if err != nil {
		return nil, err
	}
	return &pb.BookResponse{Message: "Book deleted successfully"}, nil
}
