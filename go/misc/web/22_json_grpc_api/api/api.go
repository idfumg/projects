package api

import (
	"context"
	"myapp/service"
	"myapp/types"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	listenAddr string
	svc        service.PriceFetcher
}

func NewServer(listenAddr string, svc service.PriceFetcher) *Server {
	return &Server{
		listenAddr: listenAddr,
		svc:        svc,
	}
}

func (s *Server) Run() {
	r := s.newRouter()
	http.ListenAndServe(s.listenAddr, r)
}

func (s *Server) newRouter() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", toHandlerFunc(s.handleFetchPrice))

	return r
}

func (s *Server) handleFetchPrice(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	ticker := r.URL.Query().Get("ticker")

	price, err := s.svc.FetchPrice(ctx, ticker)
	if err != nil {
		return err
	}

	out := types.PriceResponse{
		Price:  price,
		Ticker: ticker,
	}

	return writeJson(w, http.StatusOK, &out)
}
