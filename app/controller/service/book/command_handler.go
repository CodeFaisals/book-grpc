package book

import (
	"faisal.com/bookProject/repository"
	pb "faisal.com/bookProject/server/proto"
	"fmt"
	"log"
)

func HandleGetBooks() (*pb.BookListResponse, error) {
	books, err := repository.GetBooks()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch books: %v", err)
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

func HandleAddBook(bookName string) (*pb.BookResponse, error) {
	log.Printf("Adding book: %s", bookName)

	err := repository.AddBook(bookName)
	if err != nil {
		return nil, fmt.Errorf("failed to add book: %v", err)
	}

	return &pb.BookResponse{
		Message: fmt.Sprintf("Book '%s' added successfully", bookName),
	}, nil
}

func HandleDeleteBook(id string) (*pb.BookResponse, error) {
	log.Printf("Deleting book with ID: %s", id)

	err := repository.DeleteBook(id)
	if err != nil {
		return nil, fmt.Errorf("failed to delete book: %v", err)
	}

	return &pb.BookResponse{
		Message: fmt.Sprintf("Book with ID '%s' deleted successfully", id),
	}, nil
}
