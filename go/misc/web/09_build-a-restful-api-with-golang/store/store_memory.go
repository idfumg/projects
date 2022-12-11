package store

import (
	"fmt"
)

type StoreMemory struct {
	books []*Book
}

func NewStoreMemory() (*StoreMemory, error) {
	store := &StoreMemory{}
	store.books = store.createBooks()
	return store, nil
}

func (s *StoreMemory) createBooks() []*Book {
	return append([]*Book{},
		&Book{ID: 1, Title: "Golang pointers", Author: "Mr. Golang", Year: "2010"},
		&Book{ID: 2, Title: "Goroutines", Author: "Mr. Goroutine", Year: "2011"},
		&Book{ID: 3, Title: "Golang routers", Author: "Mr. Router", Year: "2012"},
		&Book{ID: 4, Title: "Golang concurrency", Author: "Mr. Currency", Year: "2013"},
		&Book{ID: 5, Title: "Golang good parts", Author: "Mr. Good", Year: "2014"})
}

func (s *StoreMemory) GetBooks() ([]*Book, error) {
	return s.books, nil
}

func (s *StoreMemory) GetBook(id int) (*Book, error) {
	for _, book := range s.books {
		if book.GetId() == id {
			return book, nil
		}
	}
	return nil, fmt.Errorf("could not find the request book with id: %d", id)
}

func (s *StoreMemory) AddBook(book *Book) (int, error) {
	s.books = append(s.books, book)
	return book.GetId(), nil
}

func (s *StoreMemory) UpdateBook(book *Book) (int, error) {
	for i := 0; i < len(s.books); i++ {
		if s.books[i].GetId() == book.GetId() {
			s.books[i] = book
			return book.GetId(), nil
		}
	}
	return 0, fmt.Errorf("book was not found")
}

func (s *StoreMemory) DeleteBook(id int) (int, error) {
	for i := 0; i < len(s.books); i++ {
		if s.books[i].GetId() == id {
			s.books = append(s.books[:i], s.books[i+1:]...)
			return id, nil
		}
	}
	return 0, fmt.Errorf("books was not found")
}
