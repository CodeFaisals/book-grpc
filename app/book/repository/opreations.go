package repository

import (
	"fmt"
	"time"

	b "github.com/BlazeCode1/book-grpc/app/book/model/Book"
	"github.com/BlazeCode1/book-grpc/couchbase"
	"github.com/couchbase/gocb/v2"
)

func GetBooks() ([]b.Book, error) {
	query := "SELECT id, book_name FROM `books_bucket`"
	rows, err := couchbase.Cluster.Query(query, nil)
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
func InsertBook(book b.Book) error {
	collection := couchbase.GetCollection()
	_, err := collection.Upsert(book.ID, book, &gocb.UpsertOptions{
		Timeout: 5 * time.Second,
	})
	if err != nil {
		return fmt.Errorf("could not insert book: %v", err)
	}
	return nil
}

func UpdateBook(id, newBookName string) error {
	_, err := couchbase.Cluster.Bucket("books_bucket").DefaultCollection().Upsert(id, map[string]interface{}{
		"id":        id,
		"book_name": newBookName,
	}, nil)

	if err != nil {
		return err
	}

	return nil
}

func DeleteBook(id string) error {
	query := "DELETE FROM `books_bucket` WHERE id = $1"
	_, err := couchbase.Cluster.Query(query, &gocb.QueryOptions{
		PositionalParameters: []interface{}{id},
	})
	return err
}
