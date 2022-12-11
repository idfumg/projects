package server

import (
	"encoding/json"
	"io"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   string `json:"year"`
}

func (b *Book) GetId() int        { return b.ID }
func (b *Book) GetTitle() string  { return b.Title }
func (b *Book) GetAuthor() string { return b.Author }
func (b *Book) GetYear() string   { return b.Year }

func NewBook(r io.Reader) (*Book, error) {
	book := &Book{}
	err := json.NewDecoder(r).Decode(book)
	return book, err
}

