package repository

import (
	"fmt"
	"time"

	b "github.com/BlazeCode1/book-grpc/app/book/model/Book"
	"github.com/BlazeCode1/book-grpc/couchbase"
	"github.com/couchbase/gocb/v2"
)

type BookRepository interface {
	GetBooks() ([]b.Book, error)
	InsertBook(book b.Book) error
	DeleteBook(id string) error
	UpdateBook(id, newBookName string) error
}

type bookRepository struct {
	couchbase couchbase.Client //Calls interface in connection.go
	//conn couchbase.Client.NewClient we gotta figure this out
}

func (s *bookRepository) GetBooks() ([]b.Book, error) {

	conn := s.couchbase.NewClient()
	query := "SELECT id, book_name,author FROM `books_bucket`"
	rows, err := conn.Cluster.Query(query, nil)
	if err != nil {
		return nil, fmt.Errorf("could not get books: %v", err)
	}

	var books []b.Book
	for rows.Next() {
		var book b.Book
		if err := rows.Row(&book); err != nil {
			return nil, fmt.Errorf("could not get book: %v", err)
		}
		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %v", err)
	}

	return books, nil
}
func (s *bookRepository) InsertBook(book b.Book) error {
	collection := s.couchbase.GetCollection()
	_, err := collection.Upsert(book.ID, book, &gocb.UpsertOptions{
		Timeout: 5 * time.Second,
	})
	if err != nil {
		return fmt.Errorf("could not insert book: %v", err)
	}
	return nil
}

func (s *bookRepository) UpdateBook(id, newBookName string) error {
	conn := s.couchbase.NewClient()
	_, err := conn.Cluster.Bucket("books_bucket").DefaultCollection().Upsert(id, map[string]interface{}{
		"id":        id,
		"book_name": newBookName,
	}, nil)

	if err != nil {
		return err
	}

	return nil
}

func (s *bookRepository) DeleteBook(id string) error {
	conn := s.couchbase.NewClient()
	query := "DELETE FROM `books_bucket` WHERE id = $1"
	_, err := conn.Cluster.Query(query, &gocb.QueryOptions{
		PositionalParameters: []interface{}{id},
	})
	return err
}
