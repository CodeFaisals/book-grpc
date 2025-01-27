package repository

import (
	"faisal.com/bookProject/couchbase"
	"fmt"
)

type Book struct {
	ID       string `json:"id"`
	BookName string `json:"book_name"`
}

func GetBooks() ([]Book, error) {
	query := "SELECT id, book_name FROM `books_bucket`"
	rows, err := couchbase.Cluster.Query(query, nil)
	if err != nil {
		return nil, fmt.Errorf("query failed: %v", err)
	}

	var books []Book
	for rows.Next() {
		var book Book
		if err := rows.Row(&book); err != nil {
			return nil, fmt.Errorf("row parsing failed: %v", err)
		}
		books = append(books, book)
	}
	return books, nil
}

func AddBook(bookName string) error {
	book := Book{
		ID:       uuid.New().String(),
		BookName: bookName,
	}

	collection := couchbase.GetCollection()
	_, err := collection.Upsert(book.ID, book, nil)
	return err
}

func DeleteBook(id string) error {
	query := "DELETE FROM `books_bucket` WHERE id = $id"
	_, err := couchbase.Cluster.Query(query, &gocb.QueryOptions{
		NamedParameters: map[string]interface{}{"id": id},
	})
	return err
}
