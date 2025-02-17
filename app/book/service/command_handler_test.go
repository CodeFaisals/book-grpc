package service

import (
	b "github.com/BlazeCode1/book-grpc/app/book/model/Book"
	mocks "github.com/BlazeCode1/book-grpc/mocks/app/book/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestHandleGetBooks_Success(t *testing.T) {
	mockRepo := new(mocks.MockBookRepository)

	books := []b.Book{
		{ID: "1", BookName: "Go Programming", Author: "John Doe"},
		{ID: "2", BookName: "Microservices with Go", Author: "Jane Smith"},
	}

	mockRepo.On("GetBooks").Return(books, nil)

	svc := NewBookService(mockRepo)
	resp, err := svc.HandleGetBooks()

	assert.NoError(t, err)
	assert.Len(t, resp.Books, 2)
	assert.Equal(t, "Go Programming", resp.Books[0].BookName)

	mockRepo.AssertExpectations(t)
}

func TestHandleAddBook_Success(t *testing.T) {
	mockRepo := new(mocks.MockBookRepository)

	book := b.Book{BookName: "Concurrency in Go", Author: "Katherine Cox"}

	mockRepo.On("InsertBook", mock.AnythingOfType("Book")).Return(nil)

	svc := NewBookService(mockRepo)

	resp, err := svc.HandleAddBook(book)

	assert.NoError(t, err)
	assert.Contains(t, resp.Message, "added successfully")

	mockRepo.AssertExpectations(t)
}
