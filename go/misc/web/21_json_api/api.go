package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	jwt "github.com/golang-jwt/jwt/v4"
)

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func permissionDenied(w http.ResponseWriter) {
	WriteJSON(w, http.StatusForbidden, ApiError{Error: "permission denied"})
}

func withJWTAuth(fn http.HandlerFunc, s Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Calling JWT auth middleware")

		tokenString := r.Header.Get("x-jwt-token")
		token, err := validateJWT(tokenString)
		if err != nil {
			permissionDenied(w)
			return
		}

		if !token.Valid {
			permissionDenied(w)
			return
		}

		userID, err := getID(r)
		if err != nil {
			log.Println("s.GetAccountByID(userID)")
			permissionDenied(w)
			return
		}

		account, err := s.GetAccountByID(userID)
		if err != nil {
			log.Println("s.GetAccountByID(userID)")
			permissionDenied(w)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		if account.Number != int64(claims["accountNumber"].(float64)) {
			log.Println("account.Number != claims[\"accountNumber\"]", account, claims)
			permissionDenied(w)
			return
		}

		fn(w, r)
	}
}

func GetSecret() []byte {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "12345"
	}
	return []byte(secret)
}

func createJWT(account *Account) (string, error) {
	claims := &jwt.MapClaims{
		"expiresAt":     15000,
		"accountNumber": account.Number,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	log.Println("GetSecret() = ", GetSecret())
	return token.SignedString(GetSecret())
}

func validateJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return GetSecret(), nil
	})
}

type ApiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string `json:"error"`
}

func makeHTTPHandleFunc(fn ApiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			// handle an error here
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

func getID(r *http.Request) (int, error) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		return id, fmt.Errorf("invalid id given %d", id)
	}
	return id, nil
}

type APIServer struct {
	listenAddr string
	store      Storage
}

func NewAPIServer(listenAddr string, store Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *APIServer) Run() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/accounts", withJWTAuth(makeHTTPHandleFunc(s.handleGetAccounts), s.store))
	r.Post("/accounts", makeHTTPHandleFunc(s.handleCreateAccount))
	r.Delete("/accounts/{id}", makeHTTPHandleFunc(s.handleDeleteAccount))
	r.Get("/accounts/{id}", withJWTAuth(makeHTTPHandleFunc(s.handleGetAccountById), s.store))
	r.Post("/transfer", makeHTTPHandleFunc(s.handleTransfer))
	log.Println("Listening on address:", s.listenAddr)
	http.ListenAndServe(s.listenAddr, r)
}

func (s *APIServer) handleGetAccounts(w http.ResponseWriter, r *http.Request) error {
	accounts, err := s.store.GetAccounts()
	if err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, accounts)
}

func (s *APIServer) handleGetAccountById(w http.ResponseWriter, r *http.Request) error {
	id, err := getID(r)
	if err != nil {
		return err
	}
	account, err := s.store.GetAccountByID(id)
	if err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, account)
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	in := &CreateAccountRequest{}
	err := json.NewDecoder(r.Body).Decode(in)
	if err != nil {
		return err
	}
	account := NewAccount(in.FirstName, in.LastName)
	id, err := s.store.CreateAccount(account)
	if err != nil {
		return err
	}
	account, err = s.store.GetAccountByID(id)
	if err != nil {
		return err
	}
	tokenString, err := createJWT(account)
	if err != nil {
		return err
	}
	out := map[string]string{
		"id": fmt.Sprintf("%d", id),
		"token": tokenString,
	}
	return WriteJSON(w, http.StatusCreated, out)
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	id, err := getID(r)
	if err != nil {
		return err
	}
	if err := s.store.DeleteAccount(id); err != nil {
		return err
	}
	return WriteJSON(w, http.StatusAccepted, map[string]int{"deleted": id})
}

func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	in := &TransferRequest{}
	if err := json.NewDecoder(r.Body).Decode(in); err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, in)
}
