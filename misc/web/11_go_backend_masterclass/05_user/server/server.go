package server

import (
	"context"
	"myapp/models"
	"myapp/store"

	"github.com/gin-gonic/gin"
)

type Store interface {
	CreateAccount(ctx context.Context, arg store.CreateAccountParams) (models.Account, error)
	GetAccount(ctx context.Context, id int64) (models.Account, error)
	GetAccountForUpdate(ctx context.Context, id int64) (models.Account, error)
	ListAccounts(ctx context.Context, arg store.ListAccountsParams) ([]models.Account, error)
	UpdateAccount(ctx context.Context, arg store.UpdateAccountParams) (models.Account, error)
	DeleteAccount(ctx context.Context, id int64) error
	CreateUser(ctx context.Context, arg store.CreateUserParams) (models.User, error)
	GetUser(ctx context.Context, username string) (models.User, error)
}

type Logger interface {
	Printf(format string, v ...any)
}

type Server struct {
	*gin.Engine
	store  Store
	logger Logger
}

func NewServer(store Store, logger Logger) (*Server, error) {
	server := &Server{
		gin.Default(),
		store,
		logger,
	}

	server.POST("/accounts", server.createAccount)
	server.GET("/accounts/:id", server.getAccount)
	server.GET("/accounts", server.listAccounts)
	server.POST("/users", server.createUser)

	return server, nil
}

func (s *Server) Serve(addr string) error {
	return s.Run(addr)
}

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
