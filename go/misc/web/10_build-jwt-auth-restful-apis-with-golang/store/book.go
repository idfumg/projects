package store

type Book struct {
	ID     int
	Title  string
	Author string
	Year   string
}

func (b *Book) GetId() int        { return b.ID }
func (b *Book) GetTitle() string  { return b.Title }
func (b *Book) GetAuthor() string { return b.Author }
func (b *Book) GetYear() string   { return b.Year }
