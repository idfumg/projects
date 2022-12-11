package rest

import (
	"io"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

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

type Config interface {
	GetIsTest() bool
}

type SessionData interface {
	LoadAndSave(next http.Handler) http.Handler
}

type Server struct {
	*chi.Mux
	store   Store
	logger  Logger
	cache   Cache
	config  Config
	session SessionData
}

func New(store Store, logger Logger, cache Cache, config Config) (*Server, error) {
	server := &Server{
		Mux:     chi.NewRouter(),
		store:   store,
		logger:  logger,
		cache:   cache,
		config:  config,
		session: createSession(config),
	}

	server.Use(middleware.Recoverer)
	server.Use(server.noSurfMiddleware)
	server.Use(server.sessionLoadMiddleware)

	server.Get("/", http.HandlerFunc(server.home))
	server.Get("/about", http.HandlerFunc(server.about))

	return server, nil
}

func (s *Server) Serve(port string) error {
	return http.ListenAndServe(port, addCORS(s))
}
