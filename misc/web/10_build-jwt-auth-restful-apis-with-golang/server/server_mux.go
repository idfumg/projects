package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

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
	Signup(email, password string) (int, error)
	Login(email_ string) (id int, email string, password string, err error)
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

	router.HandleFunc("/signup", server.signup).Methods("POST")
	router.HandleFunc("/login", server.login).Methods("POST")
	router.HandleFunc("/protected-endpoint", server.TokenVerifyMiddleWare(server.protectedEndpoint)).Methods("POST")

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
		ID:     book.GetId(),
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

func (s *ServerMux) signup(w http.ResponseWriter, r *http.Request) {
	s.logger.Printf("API. signup was invoked")

	user := &User{}
	json.NewDecoder(r.Body).Decode(user)

	if user.Email == "" || user.Password == "" {
		makeResponse(w, http.StatusBadRequest, NewError(
			"Wrong credentials",
			ErrWrongCredentials,
		))
		return
	}

	hash, err := user.GenerateHash()
	if err != nil {
		makeResponse(w, http.StatusBadRequest, NewError(
			"Password's hash wasn't generated",
			ErrWrongCredentials,
		))
		return
	}
	user.Password = string(hash)

	id, err := s.store.Signup(user.Email, user.Password)
	if err != nil {
		makeResponse(w, http.StatusInternalServerError, NewError(
			fmt.Sprintf("The user wasn't signed up: %v\n", err),
			ErrWrongCredentials,
		))
		return
	}

	user.ID = id
	json.NewEncoder(w).Encode(user)
}

func (s *ServerMux) login(w http.ResponseWriter, r *http.Request) {
	s.logger.Printf("API. login was invoked")

	user := &User{}
	json.NewDecoder(r.Body).Decode(user)

	if user.Email == "" || user.Password == "" {
		makeResponse(w, http.StatusBadRequest, NewError(
			"Wrong credentials",
			ErrWrongCredentials,
		))
		return
	}

	id, _, hash, err := s.store.Login(user.Email)
	if err != nil {
		makeResponse(w, http.StatusInternalServerError, NewError(
			fmt.Sprintf("The user wasn't found: %v\n", err),
			ErrWrongCredentials,
		))
		return
	}

	err = user.CheckPassword(hash)
	if err != nil {
		makeResponse(w, http.StatusInternalServerError, NewError(
			fmt.Sprintf("The password is invalid: %v\n", err),
			ErrWrongCredentials,
		))
		return
	}

	user.ID = id
	user.Password = hash
	token, err := user.GenerateToken()
	if err != nil {
		makeResponse(w, http.StatusInternalServerError, NewError(
			fmt.Sprintf("The token wasn't generated: %v\n", err),
			ErrWrongCredentials,
		))
		return
	}

	resp := struct {
		Token string `json:"token"`
	}{
		Token: token,
	}

	makeResponse(w, http.StatusOK, resp)
}

func (s *ServerMux) TokenVerifyMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.logger.Printf("API. TokenVerifyMiddleWare was invoked")

		authHeader := r.Header.Get("Authorization")
		splitted := strings.Split(authHeader, " ")
		if len(splitted) != 2 {
			makeResponse(w, http.StatusUnauthorized, NewError(
				"The valid token wasn't provided 1",
				ErrWrongCredentials,
			))
			return
		}
		token := splitted[1]

		if !IsTokenValid(token) {
			makeResponse(w, http.StatusUnauthorized, NewError(
				"The valid token wasn't provided 2",
				ErrWrongCredentials,
			))
			return
		}

		next(w, r)
	})
}

func (s *ServerMux) protectedEndpoint(w http.ResponseWriter, r *http.Request) {
	s.logger.Printf("API. protectedEndpoint was invoked")

	ans := struct {
		Message string `json:"message"`
	}{
		Message: "Token was validated. Everything is fine!",
	}
	json.NewEncoder(w).Encode(ans)
}
