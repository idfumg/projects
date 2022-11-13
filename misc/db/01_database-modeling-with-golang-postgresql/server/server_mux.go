package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"myapp/store"
)

type Store interface {
	GetBooks() ([]*store.Book, error)
	GetBook(int) (*store.Book, error)
	AddBook(*store.Book) (int, error)
	UpdateBook(*store.Book) (int, error)
	DeleteBook(id int) (int, error)
}

type Logger interface {
	Printf(format string, v ...any)
}

type ServerMux struct {
	*mux.Router
	store  Store
	logger Logger
}

func NewServerMux(store Store, logger Logger) (*ServerMux, error) {
	router := mux.NewRouter()

	server := &ServerMux{
		Router: router,
		store:  store,
		logger: logger,
	}

	router.HandleFunc("/books", server.getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", server.getBook).Methods("GET")
	router.HandleFunc("/books", server.addBook).Methods("POST")
	router.HandleFunc("/books", server.updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", server.deleteBooks).Methods("DELETE")

	return server, nil
}

func AddCORS(s *ServerMux) http.Handler {
	return handlers.CORS(
		handlers.AllowedHeaders(
			[]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods(
			[]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*"}))(s)
}

func (s *ServerMux) getBooks(w http.ResponseWriter, r *http.Request) {
	s.logger.Printf("API. getBooks was invoked")

	resp, err := s.store.GetBooks()
	if err != nil {
		makeResponse(w, http.StatusInternalServerError, NewError(
			fmt.Sprintf("Books weren't found: %v", err),
			ErrBooksWereNotFound,
		))
		return
	}

	makeResponse(w, http.StatusOK, resp)
}

func (s *ServerMux) getBook(w http.ResponseWriter, r *http.Request) {
	s.logger.Printf("API. getBook was invoked")

	id, err := getBookIdFromReq(r)
	if err != nil {
		makeResponse(w, http.StatusBadRequest, NewError(
			fmt.Sprintf("Id wasn't parsed: %v", err),
			ErrIdWasNotParsed,
		))
		return
	}

	resp, err := s.store.GetBook(id)
	if err != nil {
		makeResponse(w, http.StatusNotFound, NewError(
			fmt.Sprintf("Book wasn't found: %v", err),
			ErrBookWasNotFound,
		))
		return
	}

	makeResponse(w, http.StatusOK, resp)
}

func (s *ServerMux) addBook(w http.ResponseWriter, r *http.Request) {
	s.logger.Printf("API. addBook was invoked")

	book, err := NewBook(r.Body)
	if err != nil {
		makeResponse(w, http.StatusBadRequest, NewError(
			fmt.Sprintf("Book wasn't parsed: %v", err),
			ErrBookWasNotParsed,
		))
		return
	}

	id, err := s.store.AddBook(&store.Book{
		ID:     0,
		Title:  book.GetTitle(),
		Author: book.GetAuthor(),
		Year:   book.GetYear(),
	})
	if err != nil {
		makeResponse(w, http.StatusInternalServerError, NewError(
			fmt.Sprintf("Book wasn't inserted: %v", err),
			ErrBookWasNotInserted,
		))
		return
	}

	resp := &struct {
		Id int `json:"id"`
	}{
		Id: id,
	}

	makeResponse(w, http.StatusCreated, resp)
}

func (s *ServerMux) updateBook(w http.ResponseWriter, r *http.Request) {
	s.logger.Printf("API. updateBook was invoked")

	book, err := NewBook(r.Body)
	if err != nil {
		makeResponse(w, http.StatusBadRequest, NewError(
			fmt.Sprintf("Book wasn't parsed: %v", err),
			ErrBookWasNotParsed,
		))
		return
	}

	id, err := s.store.UpdateBook(&store.Book{
		ID:     book.GetId(),
		Title:  book.GetTitle(),
		Author: book.GetAuthor(),
		Year:   book.GetYear(),
	})

	if err != nil {
		makeResponse(w, http.StatusBadRequest, NewError(
			fmt.Sprintf("Book wasn't found: %v", err),
			ErrBookWasNotFound,
		))
		return
	}

	resp := &struct {
		Id int `json:"id"`
	}{
		Id: id,
	}

	makeResponse(w, http.StatusAccepted, resp)
}

func (s *ServerMux) deleteBooks(w http.ResponseWriter, r *http.Request) {
	s.logger.Printf("API. deleteBooks was invoked")

	id, err := getBookIdFromReq(r)
	if err != nil {
		makeResponse(w, http.StatusBadRequest, NewError(
			fmt.Sprintf("Id wasn't parsed: %v", err),
			ErrIdWasNotParsed,
		))
		return
	}

	id, err = s.store.DeleteBook(id)
	if err != nil {
		makeResponse(w, http.StatusBadRequest, NewError(
			fmt.Sprintf("Book wasn't found: %v", err),
			ErrBookWasNotFound,
		))
		return
	}

	resp := &struct {
		Id int `json:"id"`
	}{
		Id: id,
	}

	makeResponse(w, http.StatusAccepted, resp)
}

func makeResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func getBookIdFromReq(r *http.Request) (int, error) {
	params := mux.Vars(r)
	return strconv.Atoi(params["id"])
}
