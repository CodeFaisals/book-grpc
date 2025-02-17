package service

import (
	b "github.com/BlazeCode1/book-grpc/app/book/model/Book"
	mocks "github.com/BlazeCode1/book-grpc/mocks2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestHandleGetBooks_Success(t *testing.T) {
	mockRepo := new(mocks.BookRepository)

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
	assert.Equal(t, "Jane Smith", resp.Books[1].Author)
	assert.Equal(t, "John Doe", resp.Books[0].Author)
	mockRepo.AssertExpectations(t)
}

//func TestHandleGetBooks_Failure(t *testing.T) {
//	mockRepo := new(mocks.BookRepository)
//	books := []b.Book{}
//
//	mockRepo.On("GetBooks").Return(books, nil)
//
//	svc := NewBookService(mockRepo)
//	resp, err := svc.HandleGetBooks()
//
//	assert.Error(t, err)
//	assert.Nil(t, resp)
//	mockRepo.AssertExpectations(t)
//}

func TestHandleAddBook_Success(t *testing.T) {
	mockRepo := new(mocks.BookRepository)

	book := b.Book{BookName: "Concurrency in Go", Author: "Katherine Cox"}

	mockRepo.On("InsertBook", mock.AnythingOfType("Book")).Return(nil)

	svc := NewBookService(mockRepo)

	resp, err := svc.HandleAddBook(book)

	assert.NoError(t, err)
	assert.Contains(t, resp.Message, "added successfully")

	mockRepo.AssertExpectations(t)
}

func TestHandleDeleteBook_Success(t *testing.T) {
	mockRepo := new(mocks.BookRepository)
	mockRepo.On("DeleteBook", mock.AnythingOfType("string")).Return(nil)
	book := b.Book{ID: "1", BookName: "Concurrency in Go", Author: "Katherine Cox"}

	svc := NewBookService(mockRepo)

	resp, err := svc.HandleDeleteBook(book.ID)
	assert.NoError(t, err)
	assert.Contains(t, resp.Message, "deleted successfully")
	mockRepo.AssertExpectations(t)
}

func TestHandleUpdateBook_Success(t *testing.T) {
	mockRepo := new(mocks.BookRepository)
	mockRepo.On("UpdateBook", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil)
	book := b.Book{ID: "1", BookName: "test", Author: "test"}
	svc := NewBookService(mockRepo)
	resp, err := svc.HandleUpdateBook(book.ID, book)

	assert.NoError(t, err)
	assert.Contains(t, resp.Message, "updated successfully")
	mockRepo.AssertExpectations(t)
}
