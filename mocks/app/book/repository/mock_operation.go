package repository

import (
	b "github.com/BlazeCode1/book-grpc/app/book/model/Book"
	"github.com/stretchr/testify/mock"
)

type MockBookRepository struct {
	mock.Mock
}

func (m *MockBookRepository) GetBooks() ([]b.Book, error) {
	args := m.Called()
	return args.Get(0).([]b.Book), args.Error(1)
}

func (m *MockBookRepository) InsertBook(book b.Book) error {
	args := m.Called(book)
	return args.Error(0)
}

func (m *MockBookRepository) UpdateBook(id, newBookName, author string) error {
	args := m.Called(id, newBookName, author)
	return args.Error(0)
}

func (m *MockBookRepository) DeleteBook(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
