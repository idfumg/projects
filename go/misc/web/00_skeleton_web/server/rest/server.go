package rest

import (
	"io"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/gorilla/handlers"
	"github.com/justinas/nosurf"

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

type Cache interface {
	Render(w io.Writer, name string, td *TemplateData) error
}

type Server struct {
	*chi.Mux
	store  Store
	logger Logger
	cache  Cache
}

func New(store Store, logger Logger, cache Cache) (*Server, error) {
	server := &Server{
		Mux:    chi.NewRouter(),
		store:  store,
		logger: logger,
		cache:  cache,
	}

	server.Use(middleware.Recoverer)
	server.Use(noSurf)

	server.Get("/", http.HandlerFunc(server.home))
	server.Get("/about", http.HandlerFunc(server.about))

	return server, nil
}

func (s *Server) Serve(port string) error {
	return http.ListenAndServe(port, addCORS(s))
}

func addCORS(s *Server) http.Handler {
	return handlers.CORS(
		handlers.AllowedHeaders(
			[]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods(
			[]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*"}))(s)
}

func noSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}
