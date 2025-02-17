package service

import (
	"fmt"
	"log"

	pb "github.com/BlazeCode1/book-grpc/app/book/controller/grpc"
	b "github.com/BlazeCode1/book-grpc/app/book/model/Book"
	operation "github.com/BlazeCode1/book-grpc/app/book/repository"
	"github.com/google/uuid"
)

type BookService interface {
	HandleGetBooks() (*pb.BookListResponse, error)
	HandleAddBook(book b.Book) (*pb.BookResponse, error)
	HandleDeleteBook(id string) (*pb.BookResponse, error)
	HandleUpdateBook(id string, book b.Book) (*pb.BookResponse, error)
}

type bookService struct {
	repo operation.BookRepository
}

func NewBookService(repo operation.BookRepository) BookService {
	return &bookService{repo: repo}
}

func (s *bookService) HandleGetBooks() (*pb.BookListResponse, error) {
	// remove this query and add it to repository/operations.go # DONE
	books, err := s.repo.GetBooks()
	if err != nil {
		return nil, fmt.Errorf("failed to get books: %v", err)
	}

	var pbBooks []*pb.Book
	for _, book := range books {
		pbBooks = append(pbBooks, &pb.Book{
			Id:       book.ID,
			BookName: book.BookName,
			Author:   book.Author,
		})
	}
	return &pb.BookListResponse{Books: pbBooks}, nil
}

func (s *bookService) HandleAddBook(book b.Book) (*pb.BookResponse, error) {
	log.Printf("Adding book: %s", book.BookName)

	// Create a new Book instance
	bookInstance := b.Book{
		ID:       uuid.New().String(), // Generate a new UUID for each book
		BookName: book.BookName,
		Author:   book.Author,
	}

	err := s.repo.InsertBook(bookInstance)
	if err != nil {
		return nil, fmt.Errorf("failed to add book: %v", err)
	}

	return &pb.BookResponse{
		Message: fmt.Sprintf("Book '%s' added successfully", bookInstance.BookName),
	}, nil
}

func (s *bookService) HandleDeleteBook(id string) (*pb.BookResponse, error) {
	log.Printf("Deleting book with ID: %s", id)
	// Delete the book from your storage
	err := s.repo.DeleteBook(id)
	if err != nil {
		return nil, err
	}
	return &pb.BookResponse{Message: "Book deleted successfully"}, nil
}

func (s *bookService) HandleUpdateBook(id string, book b.Book) (*pb.BookResponse, error) {
	log.Printf("Updating book with ID: %s", id)
	err := s.repo.UpdateBook(id, book.BookName, book.Author)
	if err != nil {
		return nil, err
	}
	return &pb.BookResponse{Message: "Book updated successfully"}, nil
}
