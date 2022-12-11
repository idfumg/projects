package store

import (
	"errors"
	"fmt"
	"time"

	pg "github.com/go-pg/pg"
	orm "github.com/go-pg/pg/orm"
)

type Logger interface {
	Printf(format string, v ...any)
}

type Config interface {
	GetDBHost() string
	GetDBPort() string
	GetDBUsername() string
	GetDBTable() string
	GetDBPassword() string
	GetDBSslMode() string
	GetDBUrl() string
}

type StorePg struct {
	db     *pg.DB
	logger Logger
	config Config
}

func NewStorePg(logger Logger, config Config) (*StorePg, func(), error) {
	opts := &pg.Options{
		User:         config.GetDBUsername(),
		Password:     config.GetDBPassword(),
		Addr:         fmt.Sprintf("%s:%s", config.GetDBHost(), config.GetDBPort()),
		Database:     config.GetDBTable(),
		DialTimeout:  30 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  2 * time.Minute,
		MaxConnAge:   1 * time.Minute,
		PoolSize:     20,
	}

	db := pg.Connect(opts)
	if db == nil {
		return nil, nil, errors.New("faild to connect to the database")
	}

	err := CeateBookTable(db)
	if err != nil {
		return nil, nil, errors.New("faild to create tables")
	}

	cancel := func() {
		logger.Printf("Closing database connection")
		db.Close()
	}

	return &StorePg{
		db:     db,
		logger: logger,
		config: config,
	}, cancel, nil
}

func CeateBookTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}

	return db.CreateTable(&Book{}, opts)
}

func (s *StorePg) GetBooks() ([]*Book, error) {
	s.logger.Printf("Store. GetBooks was invoked")

	var books []*Book
	err := s.db.Model(&books).Select()
	return books, err
}

func (s *StorePg) GetBook(id int) (*Book, error) {
	s.logger.Printf("Store. GetBook was invoked")

	book := &Book{ID: id}
	err := s.db.Model(book).WherePK().Select()
	return book, err
}

func (s *StorePg) AddBook(book *Book) (int, error) {
	s.logger.Printf("Store. AddBook was invoked")

	err := s.db.Insert(book)
	return book.GetId(), err
}

func (s *StorePg) UpdateBook(book *Book) (int, error) {
	s.logger.Printf("Store. UpdateBook was invoked")

	_, err := s.db.Model(book).WherePK().Update()
	return book.GetId(), err
}

func (s *StorePg) DeleteBook(id int) (int, error) {
	s.logger.Printf("Store. DeleteBook was invoked")

	book := &Book{ID: id}
	_, err := s.db.Model(book).WherePK().Delete()
	return id, err
}
