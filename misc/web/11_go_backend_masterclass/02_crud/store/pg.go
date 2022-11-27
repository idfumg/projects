package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
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

type DbTx interface {
	ExecContext(ctx context.Context, query string, params ...interface{}) (sql.Result, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	QueryContext(ctx context.Context, query string, params ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, params ...interface{}) *sql.Row
}

type Pg struct {
	db     DbTx
	logger Logger
	config Config
}

func NewPg(logger Logger, config Config) (*Pg, error) {
	connectString, err := GetConnectString(config)
	if err != nil {
		return nil, err
	}

	db, err := sqlx.Connect("postgres", connectString)
	if err != nil {
		return nil, errors.New("could not connect to the database: " + err.Error())
	}

	if err := db.PingContext(context.Background()); err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(200)
	db.SetMaxIdleConns(0)
	db.SetConnMaxLifetime(30 * time.Second)
	db.SetConnMaxIdleTime(30 * time.Second)

	return &Pg{
		db:     db,
		logger: logger,
		config: config,
	}, nil
}

func GetConnectString(config Config) (string, error) {
	if len(config.GetDBHost()) != 0 {
		return GetDBUrl(
			config.GetDBHost(),
			config.GetDBPort(),
			config.GetDBUsername(),
			config.GetDBTable(),
			config.GetDBPassword(),
			config.GetDBSslMode()), nil
	}
	return pq.ParseURL(config.GetDBUrl()) // postgres://postgres:postgres@localhost:5432/dbname?sslmode=disable
}

func GetDBUrl(host, port, username, table, password, sslmode string) string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		host, port, username, table, password, sslmode,
	)
}

func (s *Pg) GetBooks() ([]*Book, error) {
	s.logger.Printf("Store. GetBooks was invoked")
	return nil, nil
	// rows, err := s.db.Query("select id, title, author, year from books")
	// if err != nil {
	// 	return nil, err
	// }

	// defer rows.Close()

	// ans := []*Book{}

	// for rows.Next() {
	// 	book := &Book{}

	// 	err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	ans = append(ans, book)
	// }

	// if rows.Err() != nil {
	// 	return nil, err
	// }

	// return ans, nil
}

func (s *Pg) GetBook(id int) (*Book, error) {
	s.logger.Printf("Store. GetBook was invoked")

	// book := &Book{}
	// row := s.db.QueryRow("select id, title, author, year from books where id = $1", id)
	// err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
	// if err != nil {
	// 	return nil, err
	// }

	return nil, nil
}

func (s *Pg) AddBook(book *Book) (int, error) {
	s.logger.Printf("Store. AddBook was invoked")
	return 0, nil
	// id := 0
	// err := s.db.QueryRow("insert into books (title, author, year) values ($1, $2, $3) RETURNING id;",
	// 	book.GetTitle(),
	// 	book.GetAuthor(),
	// 	book.GetYear()).Scan(&id)

	// return id, err
}

func (s *Pg) UpdateBook(book *Book) (int, error) {
	s.logger.Printf("Store. UpdateBook was invoked")
	return 0, nil
	// res, err := s.db.Exec("update books set title=$1, author=$2, year=$3 where id=$4 returning id;",
	// 	book.GetTitle(), book.GetAuthor(), book.GetYear(), book.GetId())

	// if err != nil {
	// 	return 0, err
	// }

	// cnt, err := res.RowsAffected()
	// if err != nil {
	// 	return 0, err
	// }
	// if cnt == 0 {
	// 	return 0, errors.New("could not find the book with the current id")
	// }

	// return book.GetId(), nil
}

func (s *Pg) DeleteBook(id int) (int, error) {
	s.logger.Printf("Store. DeleteBook was invoked")
	return 0, nil
	// res, err := s.db.Exec("delete from books where id=$1", id)
	// if err != nil {
	// 	return 0, err
	// }

	// cnt, err := res.RowsAffected()
	// if err != nil {
	// 	return 0, err
	// }
	// if cnt == 0 {
	// 	return 0, errors.New("could not find the book with the current id")
	// }

	// return id, nil
}
