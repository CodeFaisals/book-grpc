package book

import (
	"fmt"
	"log"

	pb "github.com/BlazeCode1/book-grpc/app/controller/grpc"
	"github.com/BlazeCode1/book-grpc/app/repository"
	"github.com/google/uuid"
)

func HandleGetBooks() (*pb.BookListResponse, error) {
	query := "SELECT id, book_name FROM `books_bucket`"
	rows, err := repository.Cluster.Query(query, nil)
	if err != nil {
		return nil, err
	}

	var books []*pb.Book
	for rows.Next() {
		var row repository.Book
		if err := rows.Row(&row); err != nil {
			return nil, err
		}
		books = append(books, &pb.Book{
			Id:       row.ID,
			BookName: row.BookName,
		})
	}

	return &pb.BookListResponse{Books: books}, nil
}

//func (s *server) GetBooks(ctx context.Context, req *pb.EmptyRequest) (*pb.BookListResponse, error) {
//	query := "SELECT id, book_name FROM `books_bucket`"
//	rows, err := couchbase.Cluster.Query(query, nil)
//	if err != nil {
//		return nil, err
//	}
//
//	var books []*pb.Book
//	for rows.Next() {
//		var row couchbase.Book
//		if err := rows.Row(&row); err != nil {
//			return nil, err
//		}
//		books = append(books, &pb.Book{
//			Id:       row.ID,
//			BookName: row.BookName,
//		})
//	}
//
//	return &pb.BookListResponse{
//		Books: books,
//	}, nil
//}

func HandleAddBook(book repository.Book) (*pb.BookResponse, error) {
	log.Printf("Adding book: %s", book.BookName)

	// Create a new Book instance
	bookInstance := repository.Book{
		ID:       uuid.New().String(), // Generate a new UUID for each book
		BookName: book.BookName,
	}

	err := repository.InsertBook(bookInstance)
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
	err := repository.DeleteBook(id)
	if err != nil {
		return nil, err
	}
	return &pb.BookResponse{Message: "Book deleted successfully"}, nil
}
