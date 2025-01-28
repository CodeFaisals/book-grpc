package repository

import (
	"fmt"
	"time"

	"github.com/BlazeCode1/book-grpc/couchbase"
	"github.com/couchbase/gocb/v2"
	"github.com/BlazeCode1/book-grpc/app/book/model/bookModel"
)



func InsertBook(book bookModel) error {
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
