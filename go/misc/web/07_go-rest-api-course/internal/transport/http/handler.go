package http

import (
	"encoding/json"
	"fmt"
	"myapp/internal/comment"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type Handler struct {
	Router  *mux.Router
	Service *comment.Service
}

type Response struct {
	Message string
	Error   string
}

func NewHandler(service *comment.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(log.Fields{
			"Method": r.Method,
			"Path":   r.URL.Path,
		}).Info("handled request")
		next.ServeHTTP(w, r)
	})
}

func BasicAuth(original func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info("Basic auth endpoint hit")
		user, pass, ok := r.BasicAuth()
		if user == "admin" && pass == "password" && ok {
			original(w, r)
		} else {
			sendErrorResponse(w, "not authorized", fmt.Errorf("not authorized"))
		}
	}
}

func validateToken(accessToken string) bool {
	var mySigningKey = []byte("missionimposible")
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there has been an error while parsing a token")
		}
		return mySigningKey, nil
	})
	if err != nil {
		return false
	}
	return token.Valid
}

func JWTAuth(original func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info("jwt authentication hit")
		authHeader := r.Header["Authorization"]
		if authHeader == nil {
			sendErrorResponse(w, "not authorized 1", fmt.Errorf("not authorized 1"))
			return
		}
		authHeaderParts := strings.Split(authHeader[0], " ")
		if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
			sendErrorResponse(w, "not authorized 2", fmt.Errorf("not authorized 2"))
			return
		}
		if validateToken(authHeaderParts[1]) {
			original(w, r)
			return
		} else {
			sendErrorResponse(w, "not authorized 3", fmt.Errorf("not authorized 3"))
			return
		}
	}
}

func (h *Handler) SetupRoutes() {
	log.Info("Setting up routes")

	h.Router = mux.NewRouter()

	h.Router.Use(LoggingMiddleware)

	h.Router.HandleFunc("/api/health", h.Health).Methods("GET")
	h.Router.HandleFunc("/api/comment", JWTAuth(h.GetAllComments)).Methods("GET")
	h.Router.HandleFunc("/api/comment", h.PostComment).Methods("POST")
	h.Router.HandleFunc("/api/comment/{id}", h.GetComment).Methods("GET")
	h.Router.HandleFunc("/api/comment/{id}", h.DeleteComment).Methods("DELETE")
	h.Router.HandleFunc("/api/comment/{id}", h.UpdateComment).Methods("PUT")
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	if err := sendOkResponse(w, Response{Message: "I'm alive!"}); err != nil {
		panic(err)
	}
}

func sendOkResponse(w http.ResponseWriter, resp interface{}) error {
	w.Header().Add("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(resp)
}

func sendErrorResponse(w http.ResponseWriter, message string, err error) {
	w.Header().Add("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(w).Encode(Response{Message: message, Error: err.Error()}); err != nil {
		panic(err)
	}
}
